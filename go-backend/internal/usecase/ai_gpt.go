package usecase

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
)

type aiGPTUsecase struct {
	repo         repository.SpotRepositoryInterface
	openaiClient openai.Client
}

var _ AIGenerateStreamInterface = (*aiGPTUsecase)(nil)
var _ embeddingGeneratorInterface = (*aiGPTUsecase)(nil)

// NewAIGPTUsecase は AIGPTUsecase の新しいインスタンスを作成します。
// baseUrl は OpenAI API のベースURLを指定します。
// APIキーは環境変数 OPENAI_API_KEY から取得されます。
func NewAIGPTUsecase(spotRepo repository.SpotRepositoryInterface, baseUrl string, apiKey string) *aiGPTUsecase {
	openaiClient := openai.NewClient(
		option.WithBaseURL(baseUrl),
		option.WithAPIKey(apiKey),
	)
	return &aiGPTUsecase{
		openaiClient: openaiClient,
		repo:         spotRepo,
	}
}

func (u *aiGPTUsecase) createEmbedding(ctx context.Context, text string) ([]float32, error) {
	resp, err := u.openaiClient.Embeddings.New(ctx, openai.EmbeddingNewParams{
		Model: openai.EmbeddingModelTextEmbeddingAda002,
		Input: openai.EmbeddingNewParamsInputUnion{OfArrayOfStrings: []string{text}},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create embedding: %w", err)
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no embedding data returned")
	}

	// float64 -> float32 変換
	result := make([]float32, len(resp.Data[0].Embedding))
	for i, v := range resp.Data[0].Embedding {
		result[i] = float32(v)
	}
	return result, nil
}

func (u *aiGPTUsecase) Event(ctx context.Context, prompt string) (io.ReadCloser, error) {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()

		// RAG: 関連情報をDBから取得するために、まずプロンプトをベクトル化
		embedding, err := u.createEmbedding(ctx, prompt)
		if err != nil {
			pw.CloseWithError(fmt.Errorf("failed to create embedding for prompt: %w", err))
			return
		}

		// ベクトル検索で関連スポット情報を取得
		spots, err := u.repo.SearchSpotsByEmbedding(ctx, embedding)
		if err != nil {
			pw.CloseWithError(fmt.Errorf("failed to search spots by embedding: %w", err))
			return
		}

		// プロンプトの構築
		systemPrompt, userPrompt, err := _buildPrompt(spots, prompt)
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
			Model: openai.ChatModelChatgpt4oLatest,
		})

		// ストリーミング応答の処理
		acc := openai.ChatCompletionAccumulator{}

		for stream.Next() {
			chunk := stream.Current()
			acc.AddChunk(chunk)

			if _, ok := acc.JustFinishedContent(); ok {
				slog.Debug("end event")
			}

			if tool, ok := acc.JustFinishedToolCall(); ok {
				// ツール呼び出しの処理（必要に応じて実装）
				slog.Info("Tool called", "name", tool.Name, "arguments", tool.Arguments)
				fmt.Fprintf(pw, "Tool called: %s with arguments %v\n", tool.Name, tool.Arguments)
			}

			if refusal, ok := acc.JustFinishedRefusal(); ok {
				// 拒否応答の処理（必要に応じて実装）
				slog.Error("Refused", "info", refusal)
			}

			if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
				pw.Write([]byte(chunk.Choices[0].Delta.Content))
			}
		}

		if err := stream.Err(); err != nil {
			pw.CloseWithError(fmt.Errorf("stream error: %w", err))
			ctx.Err()
			return
		}
	}()

	return pr, nil
}
