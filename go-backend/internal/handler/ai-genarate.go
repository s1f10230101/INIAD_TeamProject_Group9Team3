package handler

import (
	"context"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// 旅行プランの生成
// (POST /plans)
func (s *server) GeneratePlan(ctx context.Context, request oapi.GeneratePlanRequestObject) (oapi.GeneratePlanResponseObject, error) {
	return oapi.GeneratePlan500JSONResponse{Message: "Not Implemented"}, nil
}
