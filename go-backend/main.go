package main

import (
	"net/http"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/handler"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/usecase"
)

func main() {
	postRepositoryInmemory := repository.NewPostRepositoryInmemory()
	postUsecase := usecase.NewPostUseCase(postRepositoryInmemory)
	serverMethods := handler.NewServer(postUsecase)

	handler := api.NewStrictHandler(serverMethods, nil)

	server := http.NewServeMux()

	api.HandlerFromMuxWithBaseURL(handler, server, "/api/v1")
	http.ListenAndServe(":8080", server)
}
