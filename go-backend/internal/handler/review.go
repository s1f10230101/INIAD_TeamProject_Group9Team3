package handler

import (
	"context"
	"log/slog"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// 観光スポットのレビュー一覧の取得
// (GET /reviews)
func (s *server) GetReviewsBySpotId(ctx context.Context, request oapi.GetReviewsBySpotIdRequestObject) (oapi.GetReviewsBySpotIdResponseObject, error) {
	reviews, err := s.reviewUC.GetReviewsBySpotID(ctx, request.SpotId)
	if err != nil {
		slog.Error("Error getting reviews by spot ID", slog.String("error", err.Error()))
		return nil, err
	}
	return oapi.GetReviewsBySpotId200JSONResponse(reviews), nil
}

// レビューの投稿
// (POST /reviews)
func (s *server) CreateReview(ctx context.Context, request oapi.CreateReviewRequestObject) (oapi.CreateReviewResponseObject, error) {
	res, err := s.reviewUC.CreateReview(ctx, request.SpotId, request.Body)
	if err != nil {
		slog.Error("Error creating review", slog.String("error", err.Error()))
		return nil, err
	}
	return oapi.CreateReview201JSONResponse(*res), nil
}
