package handler

import (
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/usecase"
)

type server struct {
	// Add server fields here
	postUC usecase.PostUseCaseInterface
}

func NewServer(postuc usecase.PostUseCaseInterface) api.StrictServerInterface {
	return &server{postUC: postuc}
}
