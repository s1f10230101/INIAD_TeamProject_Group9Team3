package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/handler"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/usecase"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

func main() {
	POSTGRES_USER, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		slog.Error("環境変数", "POSTGRESUSER", ok)
	}
	POSTGRES_PASSWORD, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		slog.Error("環境変数", "POSTGRES_PASSWORD", ok)
	}
	POSTGRES_DB, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		slog.Error("環境変数", "POSTGRES_DB", ok)
	}
	DBHOST, ok := os.LookupEnv("DB_HOST")
	if !ok {
		slog.Error("環境変数", "DBHOST", ok)
	}
	OPENAI_API_BASE, ok := os.LookupEnv("OPENAI_API_BASE")
	if !ok {
		slog.Error("環境変数", "OPENAI_API_BASE", OPENAI_API_BASE)
	}
	OPENAI_API_KEY, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		slog.Error("環境変数", "OPENAI_API_KEY", OPENAI_API_KEY)
	}
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=enable",
		POSTGRES_USER,
		POSTGRES_PASSWORD,
		DBHOST,
		5432,
		POSTGRES_DB)
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	m, err := migrate.New(
		"file:///./app/db/migration", dbURL,
	)
	if err != nil {
		slog.Error("migration", "migration err", err.Error())
	}
	if err := m.Up(); err != nil {
		slog.Error("migration", "migration err", err.Error())
	}
	defer m.Close()
	defer pool.Close()

	// 1. レポジトリの初期化 (Postgres版を使用)
	spotRepository := repository.NewPostgresPostRepository(pool)
	reviewRepository := repository.NewPostgresReviewRepository(pool)

	// 2. ユースケースのインスタンスを作成し、レポジトリを注入
	aiUsecase := usecase.NewAIGPTUsecase(spotRepository, OPENAI_API_BASE, OPENAI_API_KEY)
	postUsecase := usecase.NewPostUseCase(spotRepository, aiUsecase)
	reviewUsecase := usecase.NewReviewUseCase(reviewRepository)

	// 5. ハンドラを作成し、ユースケースを注入
	serverMethods := handler.NewServer(postUsecase, reviewUsecase, aiUsecase)
	handlerFuncs := oapi.NewStrictHandler(serverMethods, []oapi.StrictMiddlewareFunc{})

	// 4. HTTPサーバーの設定と起動(標準ライブラリのnet/httpを使用)
	server := oapi.HandlerWithOptions(handlerFuncs, oapi.StdHTTPServerOptions{
		BaseURL:    "/v1",
		BaseRouter: http.NewServeMux(),
		Middlewares: []oapi.MiddlewareFunc{
			handler.LoggingMiddleware,
			//handler.CorsMiddleware,
		},
	})

	// 5. ハンドラをサーバーに登録
	log.Println("Server is running on http://localhost:8080/v1")
	//if err := http.ListenAndServe(":8080", server); err != nil {
	if err := http.ListenAndServe("0.0.0.0:8080", handler.CorsMiddleware(server)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
