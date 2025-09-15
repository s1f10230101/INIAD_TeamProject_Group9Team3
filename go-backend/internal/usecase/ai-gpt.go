package usecase

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"text/template"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

type AIGPTUsecase struct {
	repo         repository.SpotRepositoryInterface
	openaiClient openai.Client
}

var _ AIGenerateStreamInterface = (*AIGPTUsecase)(nil)

// NewAIGPTUsecase は AIGPTUsecase の新しいインスタンスを作成します。
// baseUrl は OpenAI API のベースURLを指定します。
// APIキーは環境変数 OPENAI_API_KEY から取得されます。
func NewAIGPTUsecase(repo repository.SpotRepositoryInterface, baseUrl string) *AIGPTUsecase {
	openaiClient := openai.NewClient(
		option.WithBaseURL(baseUrl),
	)
	return &AIGPTUsecase{
		repo:         repo,
		openaiClient: openaiClient,
	}
}

func buildPrompt(spots []oapi.SpotResponse, userPromptInput string) (string, string, error) {
	var systemPrompt strings.Builder
	var userPrompt strings.Builder
	// 取得した情報をプロンプトに組み込む
	systemPromptText := `
あなたは旅行プランのプロです。以下の参考情報とユーザーの要望を元に、魅力的な旅行プランを提案してください。
参考情報:
{{- range .Spots }}
- 名前: {{ .Name }}
  説明: {{ .Description }}
  住所: {{ .Address }}
{{- end }}
`
	dataMap := map[string]interface{}{
		"Spots": spots,
	}
	ts := template.Must(template.New("systemPrompt").Parse(systemPromptText))
	if err := ts.Execute(&systemPrompt, dataMap); err != nil {
		return "", "", fmt.Errorf("failed to execute system prompt template: %w", err)
	}

	userPromptText := `
{{ .UserPrompt }}
 `

	tu := template.Must(template.New("userPrompt").Parse(userPromptText))
	if err := tu.Execute(&userPrompt, map[string]interface{}{
		"UserPrompt": userPromptInput,
	}); err != nil {
		return "", "", fmt.Errorf("failed to execute user prompt template: %w", err)
	}
	return systemPrompt.String(), userPrompt.String(), nil
}

func (u *AIGPTUsecase) Event(ctx context.Context, prompt string) (io.ReadCloser, error) {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()

		// RAG: 関連情報をDBから取得
		spots, err := u.repo.SearchSpots(ctx, prompt)
		if err != nil {
			pw.CloseWithError(fmt.Errorf("failed to search spots: %w", err))
			return
		}

		// プロンプトの構築
		systemPrompt, userPrompt, err := buildPrompt(spots, prompt)
		if err != nil {
			pw.CloseWithError(fmt.Errorf("failed to build prompt: %w", err))
			slog.Error("failed to build prompt", "error", err)
			return
		}

		// OpenAI API呼び出し
		stream := u.openaiClient.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(systemPrompt),
				openai.UserMessage(userPrompt),
			},
			Model: openai.ChatModelGPT5Nano,
		})

		// ストリーミング応答の処理
		acc := openai.ChatCompletionAccumulator{}

		for stream.Next() {
			chunk := stream.Current()
			acc.AddChunk(chunk)

			if content, ok := acc.JustFinishedContent(); ok {
				pw.Write([]byte(content))
				pw.Write([]byte("end\n"))
			}

			if tool, ok := acc.JustFinishedToolCall(); ok {
				// ツール呼び出しの処理（必要に応じて実装）
				fmt.Printf("Tool called: %s with arguments %v\n", tool.Name, tool.Arguments)
			}

			if refusal, ok := acc.JustFinishedRefusal(); ok {
				// 拒否応答の処理（必要に応じて実装）
				fmt.Printf("Refusal: %s\n", refusal)
			}
			// JustFinishedイベントを処理した後にchunksを使う
			if len(chunk.Choices) > 0 {
				fmt.Printf("Chunks: %+v\n", chunk.Choices[0].Delta)
			}
		}
	}()

	return pr, nil
}
