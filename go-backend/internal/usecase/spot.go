package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
)

// handler から呼び出されるビジネスロジックのインターフェース
type PostUseCaseInterface interface {
	GetAllSpots(ctx context.Context) ([]api.Spot, error)
	CreateSpot(ctx context.Context, spot *api.SpotInput) (api.Spot, error)
	GetSpotByID(ctx context.Context, spotId uuid.UUID) (api.Spot, error)
	UpdateSpotByID(ctx context.Context, spotId uuid.UUID, spot *api.SpotInput) (api.Spot, error)
}

type postUseCase struct {
	repository repository.SpotRepositoryInterface
}

func NewPostUseCase(repo repository.SpotRepositoryInterface) *postUseCase {
	return &postUseCase{
		repository: repo,
	}
}

func (u *postUseCase) GetAllSpots(ctx context.Context) ([]api.Spot, error) {
	return u.repository.GetAllSpots(ctx)
}

func (u *postUseCase) CreateSpot(ctx context.Context, spot *api.SpotInput) (api.Spot, error) {
	return u.repository.CreateSpot(ctx, spot)
}

func (u *postUseCase) GetSpotByID(ctx context.Context, spotId uuid.UUID) (api.Spot, error) {
	return u.repository.GetSpotByID(ctx, spotId)
}

func (u *postUseCase) UpdateSpotByID(ctx context.Context, spotId uuid.UUID, spot *api.SpotInput) (api.Spot, error) {
	return u.repository.UpdateSpotByID(ctx, spotId, spot)
}
