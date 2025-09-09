package handler_test

import (
	"net/http"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/handler"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/usecase"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

func setupTestRouter() http.Handler {
	// 1. レポジトリの初期化
	spotRepositoryInmemory := repository.NewSpotRepositoryInmemory()
	reviewRepositoryInmemory := repository.NewReviewRepositoryInmemory()
	// 2. ユースケースのインスタンスを作成し、レポジトリを注入
	postUsecase := usecase.NewPostUseCase(spotRepositoryInmemory)
	reviewUsecase := usecase.NewReviewUseCase(reviewRepositoryInmemory)
	// 3. ハンドラを作成し、ユースケースを注入
	serverMethods := handler.NewServer(postUsecase, reviewUsecase)
	handler := oapi.NewStrictHandler(serverMethods, nil)
	// 4. HTTPサーバーの設定と起動(標準ライブラリのnet/httpを使用)
	server := http.NewServeMux()
	// 5. ハンドラをサーバーに登録
	oapi.HandlerFromMuxWithBaseURL(handler, server, "/v1")
	return server
}
