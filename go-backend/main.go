package main

import (
	"log"
	"net/http"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/handler"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/usecase"
)

func main() {
	// 1. レポジトリの初期化
	postRepositoryInmemory := repository.NewPostRepositoryInmemory()

	// 2. ユースケースのインスタンスを作成し、レポジトリを注入
	postUsecase := usecase.NewPostUseCase(postRepositoryInmemory)

	// 3. ハンドラを作成し、ユースケースを注入
	serverMethods := handler.NewServer(postUsecase)
	handler := api.NewStrictHandler(serverMethods, nil)

	// 4. HTTPサーバーの設定と起動(標準ライブラリのnet/httpを使用)
	server := http.NewServeMux()

	// 5. ハンドラをサーバーに登録
	api.HandlerFromMuxWithBaseURL(handler, server, "/v1")
	http.ListenAndServe(":8080", server)
	log.Println("Server is running on http://localhost:8080/v1")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
