package usecase

import (
	"context"
	"fmt"
	"io"
	"strings"
	"text/template"
	"time"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

type AIGenerateStreamInterface interface {
	// 旅行プランをストリーミング形式で生成、行単位で返す
	Event(ctx context.Context, prompt string) (io.ReadCloser, error)
}

// EmbeddingGeneratorInterface はテキストから埋め込みベクトルを生成するインターフェース
type embeddingGeneratorInterface interface {
	// CreateEmbedding は与えられたテキストから埋め込みベクトルを生成します。
	createEmbedding(ctx context.Context, text string) ([]float32, error)
}

type AIGenerateFake struct {
}

var _ AIGenerateStreamInterface = (*AIGenerateFake)(nil)
var _ embeddingGeneratorInterface = (*AIGenerateFake)(nil)

func NewAIGenerateFake() *AIGenerateFake {
	return &AIGenerateFake{}
}

// Event はストリーミング形式で旅行プランを生成するサンプル実装
func (a *AIGenerateFake) Event(ctx context.Context, prompt string) (io.ReadCloser, error) {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		for i := range 3 {
			sseMessage := fmt.Sprintf("これはテストメッセージ カウント %d", i+1)
			_, err := fmt.Fprintln(pw, sseMessage)
			if err != nil {
				fmt.Printf("Error writing to pipe: %v\n", err)
				return
			}
			time.Sleep(100 * time.Millisecond)
			select {
			case <-ctx.Done():
				fmt.Println("Context cancelled, stopping message generation")
				return
			default:
			}
		}
		// 最後にプロンプトを返す
		_, err := fmt.Fprintf(pw, "prompt: %s\n", prompt)
		if err != nil {
			fmt.Printf("Error writing prompt to pipe: %v\n", err)
			return
		}
	}()
	return pr, nil
}

// createEmbedding は与えられたテキストから埋め込みベクトルを生成するサンプル実装
func (a *AIGenerateFake) createEmbedding(ctx context.Context, text string) ([]float32, error) {
	// ダミーのベクトルを返す
	vector := make([]float32, 1536)
	for i := range vector {
		vector[i] = float32(i) * 0.001
	}
	return vector, nil
}

// buildPrompt システムプロンプトとユーザープロンプトを構築するユーティリティ関数
func _buildPrompt(spots []oapi.SpotResponse, userPromptInput string) (string, string, error) {
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
