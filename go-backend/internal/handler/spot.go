package handler

import (
	"context"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
)

func (s *server) GetSpots(ctx context.Context, request api.GetSpotsRequestObject) (api.GetSpotsResponseObject, error) {
	spots, err := s.postUC.GetAllSpots(ctx)
	return api.GetSpots200JSONResponse(spots), err
}

func (s *server) PostSpots(ctx context.Context, request api.PostSpotsRequestObject) (api.PostSpotsResponseObject, error) {
	createdSpot, err := s.postUC.CreateSpot(ctx, request.Body)
	if err != nil {
		return nil, err
	}
	return api.PostSpots201JSONResponse(createdSpot), err
}

func (s *server) GetSpotsSpotId(ctx context.Context, request api.GetSpotsSpotIdRequestObject) (api.GetSpotsSpotIdResponseObject, error) {
	spotId := request.SpotId
	spot, err := s.postUC.GetSpotByID(ctx, spotId)
	if err != nil {
		return nil, err
	}
	return api.GetSpotsSpotId200JSONResponse(spot), nil
}

func (s *server) PutSpotsSpotId(ctx context.Context, request api.PutSpotsSpotIdRequestObject) (api.PutSpotsSpotIdResponseObject, error) {
	spotId := request.SpotId
	spot, err := s.postUC.UpdateSpotByID(ctx, spotId, request.Body)
	if err != nil {
		return nil, err
	}
	return api.PutSpotsSpotId200JSONResponse(spot), nil
}
