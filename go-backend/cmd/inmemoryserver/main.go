package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/handler"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/usecase"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

func main() {
	// 1. レポジトリの初期化 (インメモリ版を使用)
	spotRepository := repository.NewSpotRepositoryInmemory()
	reviewRepository := repository.NewReviewRepositoryInmemory()

	// 2. ユースケースのインスタンスを作成し、レポジトリを注入
	fakeAiCase := usecase.NewAIGenerateFake()
	postUsecase := usecase.NewPostUseCase(spotRepository, fakeAiCase)
	reviewUsecase := usecase.NewReviewUseCase(reviewRepository)

	// 3. ハンドラを作成し、ユースケースを注入
	serverMethods := handler.NewServer(postUsecase, reviewUsecase, fakeAiCase)
	handlerFuncs := oapi.NewStrictHandler(serverMethods, []oapi.StrictMiddlewareFunc{})

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
