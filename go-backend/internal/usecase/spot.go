package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// handler から呼び出されるビジネスロジックのインターフェース
type PostUseCaseInterface interface {
	GetAllSpots(ctx context.Context) ([]oapi.SpotResponse, error)
	CreateSpot(ctx context.Context, spot *oapi.SpotResister) (oapi.SpotResponse, error)
	GetSpotByID(ctx context.Context, spotId uuid.UUID) (oapi.SpotResponse, error)
	UpdateSpotByID(ctx context.Context, spotId uuid.UUID, spot *oapi.SpotUpdate) (oapi.SpotResponse, error)
}

type postUseCase struct {
	repository repository.SpotRepositoryInterface
}

var _ PostUseCaseInterface = (*postUseCase)(nil)

func NewPostUseCase(repo repository.SpotRepositoryInterface) *postUseCase {
	return &postUseCase{
		repository: repo,
	}
}

func (u *postUseCase) GetAllSpots(ctx context.Context) ([]oapi.SpotResponse, error) {
	return u.repository.GetAllSpots(ctx)
}

func (u *postUseCase) CreateSpot(ctx context.Context, spot *oapi.SpotResister) (oapi.SpotResponse, error) {
	return u.repository.CreateSpot(ctx, spot)
}

func (u *postUseCase) GetSpotByID(ctx context.Context, spotId uuid.UUID) (oapi.SpotResponse, error) {
	return u.repository.GetSpotByID(ctx, spotId)
}

func (u *postUseCase) UpdateSpotByID(ctx context.Context, spotId uuid.UUID, spot *oapi.SpotUpdate) (oapi.SpotResponse, error) {
	return u.repository.UpdateSpotByID(ctx, spotId, spot)
}
