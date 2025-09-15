package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/handler"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/usecase"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

func main() {
	openaiAPIKey := os.Getenv("OPENAI_API_KEY")
	if openaiAPIKey == "" {
		log.Fatalf("OPENAI_API_KEY is not set")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost" // デフォルト値
	}
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"app_user",
		"password",
		dbHost,
		5432,
		"app_db")
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()
	mg, err := migrate.New(
		"file://./db/migration",
		dbURL,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	} else {
		mg.Up()
	}

	// 1. レポジトリの初期化 (Postgres版を使用)
	spotRepository := repository.NewPostgresPostRepository(pool)
	reviewRepository := repository.NewPostgresReviewRepository(pool)

	// 2. ユースケースのインスタンスを作成し、レポジトリを注入
	openaiBaseUrl := os.Getenv("OPENAI_API_BASE")
	aiUsecase := usecase.NewAIGPTUsecase(spotRepository, openaiBaseUrl, openaiAPIKey)
	postUsecase := usecase.NewPostUseCase(spotRepository, aiUsecase)
	reviewUsecase := usecase.NewReviewUseCase(reviewRepository)

	// 5. ハンドラを作成し、ユースケースを注入
	serverMethods := handler.NewServer(postUsecase, reviewUsecase, aiUsecase)
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
