package handler

import (
	"context"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
)

func (s *server) PostPlansGenerate(ctx context.Context, request api.PostPlansGenerateRequestObject) (api.PostPlansGenerateResponseObject, error)
