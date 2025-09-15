package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/usecase"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

type server struct {
	// ユースケースのインターフェースをフィールドとして持つ
	postUC   usecase.SpotUseCaseInterface
	reviewUC usecase.ReviewUseCaseInterface
	aiUC     usecase.AIGenerateStreamInterface
}

var _ oapi.StrictServerInterface = (*server)(nil)

// NewServer は server 構造体のポインタを返すコンストラクタ関数
func NewServer(postuc usecase.SpotUseCaseInterface, reviewuc usecase.ReviewUseCaseInterface, aiuc usecase.AIGenerateStreamInterface) *server {
	return &server{
		postUC:   postuc,
		reviewUC: reviewuc,
		aiUC:     aiuc,
	}
}

func (s *server) HealthCheckOpenAPI(ctx context.Context, request oapi.HealthCheckOpenAPIRequestObject) (oapi.HealthCheckOpenAPIResponseObject, error) {
	return oapi.HealthCheckOpenAPI200TextResponse("API is running"), nil
}

// ミドルウェアアクセスログ
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request", "method", r.Method, "path", r.URL.Path, "body", r.Body)
		next.ServeHTTP(w, r)
	})
}

// CorsMiddleware はCORSを許可するミドルウェア
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
