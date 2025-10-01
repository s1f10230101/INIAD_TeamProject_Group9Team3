package handler

import (
	"context"
	"log/slog"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// 観光スポット一覧の取得
// (GET /spots)
func (s *server) GetAllSpots(ctx context.Context, request oapi.GetAllSpotsRequestObject) (oapi.GetAllSpotsResponseObject, error) {
	spots, err := s.spotUC.GetAllSpots(ctx)
	if err != nil {
		slog.Error("failed to get all spots", "err", err)
		return oapi.GetAllSpots500JSONResponse{Message: "Internal Server Error"}, nil
	}
	return oapi.GetAllSpots200JSONResponse(spots), nil
}

// 観光スポットの登録
// (POST /spots)
func (s *server) CreateSpot(ctx context.Context, request oapi.CreateSpotRequestObject) (oapi.CreateSpotResponseObject, error) {
	createdSpot, err := s.spotUC.CreateSpot(ctx, request.Body)
	if err != nil {
		slog.Error("failed to create spot", "err", err)
		return nil, err
	}
	return oapi.CreateSpot201JSONResponse(createdSpot), nil
}

// 観光スポットの詳細取得
// (GET /spots/{spotId})
func (s *server) GetSpotById(ctx context.Context, request oapi.GetSpotByIdRequestObject) (oapi.GetSpotByIdResponseObject, error) {
	spot, err := s.spotUC.GetSpotByID(ctx, request.SpotId)
	if err != nil {
		slog.Error("failed to get spot by ID", "err", err)
		return nil, err
	}
	return oapi.GetSpotById200JSONResponse(spot), nil
}

// 観光スポットの更新
// (PUT /spots/{spotId})
func (s *server) UpdateSpot(ctx context.Context, request oapi.UpdateSpotRequestObject) (oapi.UpdateSpotResponseObject, error) {
	spot, err := s.spotUC.UpdateSpotByID(ctx, request.SpotId, request.Body)
	if err != nil {
		slog.Error("failed to update spot by ID", "err", err)
		return nil, err
	}
	return oapi.UpdateSpot200JSONResponse(spot), nil
}
