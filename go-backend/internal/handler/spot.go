package handler

import (
	"context"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// 観光スポット一覧の取得
// (GET /spots)
func (s *server) GetAllSpots(ctx context.Context, request oapi.GetAllSpotsRequestObject) (oapi.GetAllSpotsResponseObject, error) {
	spots, err := s.postUC.GetAllSpots(ctx)
	return oapi.GetAllSpots200JSONResponse(spots), err
}

// 観光スポットの登録
// (POST /spots)
func (s *server) CreateSpot(ctx context.Context, request oapi.CreateSpotRequestObject) (oapi.CreateSpotResponseObject, error) {
	createdSpot, err := s.postUC.CreateSpot(ctx, request.Body)
	if err != nil {
		return nil, err
	}
	return oapi.CreateSpot201JSONResponse(createdSpot), nil
}

// 観光スポットの詳細取得
// (GET /spots/{spotId})
func (s *server) GetSpotById(ctx context.Context, request oapi.GetSpotByIdRequestObject) (oapi.GetSpotByIdResponseObject, error) {
	spot, err := s.postUC.GetSpotByID(ctx, request.SpotId)
	if err != nil {
		return nil, err
	}
	return oapi.GetSpotById200JSONResponse(spot), nil
}

// 観光スポットの更新
// (PUT /spots/{spotId})
func (s *server) UpdateSpot(ctx context.Context, request oapi.UpdateSpotRequestObject) (oapi.UpdateSpotResponseObject, error) {
	spot, err := s.postUC.UpdateSpotByID(ctx, request.SpotId, request.Body)
	if err != nil {
		return nil, err
	}
	return oapi.UpdateSpot200JSONResponse(spot), nil
}
