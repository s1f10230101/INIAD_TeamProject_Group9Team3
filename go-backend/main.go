package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/handler"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/usecase"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

func main() {
	// 環境変数からDBホストを取得、なければlocalhostをデフォルト値とする
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbURL := fmt.Sprintf("postgres://app_user:password@%s:5432/app_db?sslmode=disable", dbHost)

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	

	// 1. レポジトリの初期化 (Postgres版を使用)
	postRepository := repository.NewPostgresPostRepository(pool)
	reviewRepository := repository.NewPostgresReviewRepository(pool)

	// 2. ユースケースのインスタンスを作成し、レポジトリを注入
	postUsecase := usecase.NewPostUseCase(postRepository)
	reviewUsecase := usecase.NewReviewUseCase(reviewRepository)

	// 3. ハンドラを作成し、ユースケースを注入
	serverMethods := handler.NewServer(postUsecase, reviewUsecase)
	handlerFuncs := oapi.NewStrictHandler(serverMethods, nil)

	// 4. HTTPサーバーの設定と起動(標準ライブラリのnet/httpを使用)
	server := oapi.HandlerWithOptions(handlerFuncs, oapi.StdHTTPServerOptions{
		BaseURL:    "/v1",
		BaseRouter: http.NewServeMux(),
		Middlewares: []oapi.MiddlewareFunc{
			handler.LoggingMiddleware,
			handler.CorsMiddleware,
		},
	})

	// 5. ハンドラをサーバーに登録
	log.Println("Server is running on http://localhost:8080/v1")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
