package handler

import (
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/usecase"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

type server struct {
	// ユースケースのインターフェースをフィールドとして持つ
	postUC   usecase.PostUseCaseInterface
	reviewUC usecase.ReviewUseCaseInterface
}

var _ oapi.StrictServerInterface = (*server)(nil)

// NewServer は server 構造体のポインタを返すコンストラクタ関数
func NewServer(postuc usecase.PostUseCaseInterface, reviewuc usecase.ReviewUseCaseInterface) *server {
	return &server{
		postUC:   postuc,
		reviewUC: reviewuc,
	}
}
