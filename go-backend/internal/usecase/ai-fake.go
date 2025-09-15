package usecase

import (
	"context"
	"fmt"
	"io"
	"time"
)

type AIGenerateStreamInterface interface {
	// 旅行プランをストリーミング形式で生成、行単位で返す
	Event(ctx context.Context, prompt string) (io.ReadCloser, error)
}

type AIGenerateFake struct {
}

var _ AIGenerateStreamInterface = (*AIGenerateFake)(nil)

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
