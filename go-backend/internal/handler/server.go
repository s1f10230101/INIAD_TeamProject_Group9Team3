package handler

import (
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/usecase"
)

type server struct {
	// ユースケースのインターフェースをフィールドとして持つ
	postUC usecase.PostUseCaseInterface
}

// NewServer は server 構造体のポインタを返すコンストラクタ関数
func NewServer(postuc usecase.PostUseCaseInterface) api.StrictServerInterface {
	return &server{postUC: postuc}
}
