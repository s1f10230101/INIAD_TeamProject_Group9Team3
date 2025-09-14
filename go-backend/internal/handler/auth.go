package handler

import (
	"context"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// ユーザーログイン
// (POST /auth/login)
func (s *server) LoginUser(ctx context.Context, request oapi.LoginUserRequestObject) (oapi.LoginUserResponseObject, error) {
	return oapi.LoginUser401JSONResponse{Message: "Not Implemented"}, nil
}

// ユーザー情報取得
// (GET /auth/me)
func (s *server) GetUserInfo(ctx context.Context, request oapi.GetUserInfoRequestObject) (oapi.GetUserInfoResponseObject, error) {
	return oapi.GetUserInfo401JSONResponse{Message: "Not Implemented"}, nil
}

// ユーザー登録
// (POST /auth/register)
func (s *server) RegisterUser(ctx context.Context, request oapi.RegisterUserRequestObject) (oapi.RegisterUserResponseObject, error) {
	return oapi.RegisterUser500JSONResponse{Message: "Not Implemented"}, nil
}
