package handler

import (
	"context"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
)

func (s *server) PostAuthLogin(ctx context.Context, request api.PostAuthLoginRequestObject) (api.PostAuthLoginResponseObject, error) {
	return api.PostAuthLogin401JSONResponse{}, nil
}

func (s *server) GetAuthMe(ctx context.Context, request api.GetAuthMeRequestObject) (api.GetAuthMeResponseObject, error) {
	return api.GetAuthMe401JSONResponse{}, nil
}

func (s *server) PostAuthRegister(ctx context.Context, request api.PostAuthRegisterRequestObject) (api.PostAuthRegisterResponseObject, error) {
	return api.PostAuthRegister201JSONResponse{}, nil
}
