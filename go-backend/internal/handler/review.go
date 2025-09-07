package handler

import (
	"context"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
)

func (s *server) GetSpotsSpotIdReviews(ctx context.Context, request api.GetSpotsSpotIdReviewsRequestObject) (api.GetSpotsSpotIdReviewsResponseObject, error) {
	reviews, err := s.reviewUC.GetReviewsBySpotID(ctx, request.SpotId)
	if err != nil {
		return nil, err
	}
	return api.GetSpotsSpotIdReviews200JSONResponse(reviews), nil
}

func (s *server) PostSpotsSpotIdReviews(ctx context.Context, request api.PostSpotsSpotIdReviewsRequestObject) (api.PostSpotsSpotIdReviewsResponseObject, error) {
	createdReview, err := s.reviewUC.CreateReview(ctx, request.SpotId, request.Body)
	if err != nil {
		return nil, err
	}
	return api.PostSpotsSpotIdReviews201JSONResponse(*createdReview), nil
}
