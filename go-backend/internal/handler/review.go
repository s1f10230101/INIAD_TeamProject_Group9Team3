package handler

import (
	"context"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
)

func (s *server) GetSpotsSpotIdReviews(ctx context.Context, request api.GetSpotsSpotIdReviewsRequestObject) (api.GetSpotsSpotIdReviewsResponseObject, error) {
	return api.GetSpotsSpotIdReviews200JSONResponse{}, nil
}

func (s *server) PostSpotsSpotIdReviews(ctx context.Context, request api.PostSpotsSpotIdReviewsRequestObject) (api.PostSpotsSpotIdReviewsResponseObject, error) {
	return api.PostSpotsSpotIdReviews201JSONResponse{}, nil
}
