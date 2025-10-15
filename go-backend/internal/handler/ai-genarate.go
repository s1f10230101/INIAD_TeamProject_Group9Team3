package handler

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// 旅行プランの生成
// (POST /plans)
func (s *server) GeneratePlan(ctx context.Context, request oapi.GeneratePlanRequestObject) (oapi.GeneratePlanResponseObject, error) {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		streamer, _ := s.aiUC.Event(ctx, request.Body.Prompt)
		defer streamer.Close()

		scanner := bufio.NewScanner(streamer)
		// ストリームから1行ずつ読み込み、SSE形式でパイプに書き込む
		for scanner.Scan() {
			line := scanner.Text()
			// SSEの "data: " プレフィックスを付与して書き込む
			_, err := fmt.Fprintf(pw, "data: {\"text\": \"%s\"}\n\n", line)
			if err != nil {
				slog.Error("Error writing to pipe", slog.String("error", err.Error()))
				break
			}
		}
		if err := scanner.Err(); err != nil {
			slog.Error("Error reading from stream", slog.String("error", err.Error()))
		}
		if n, err := pw.Write([]byte("event: done\n\n")); err != nil || n == 0 {
			slog.Error("Error writing done event to pipe", slog.String("error", err.Error()))
		}
	}()
	ctx.Done()
	return oapi.GeneratePlan200TexteventStreamResponse{
		Body: pr,
		Headers: oapi.GeneratePlan200ResponseHeaders{
			CacheControl: "no-cache",
			Connection:   "keep-alive",
		},
	}, nil
}
