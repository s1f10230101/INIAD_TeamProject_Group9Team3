package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// handler から呼び出されるビジネスロジックのインターフェース
type SpotUseCaseInterface interface {
	GetAllSpots(ctx context.Context) ([]oapi.SpotResponse, error)
	CreateSpot(ctx context.Context, spot *oapi.SpotResister) (oapi.SpotResponse, error)
	GetSpotByID(ctx context.Context, spotId uuid.UUID) (oapi.SpotResponse, error)
	UpdateSpotByID(ctx context.Context, spotId uuid.UUID, spot *oapi.SpotUpdate) (oapi.SpotResponse, error)
}

type spotUseCase struct {
	repository   repository.SpotRepositoryInterface
	aiEmmbedding embeddingGeneratorInterface
}

var _ SpotUseCaseInterface = (*spotUseCase)(nil)

func NewPostUseCase(repo repository.SpotRepositoryInterface, aiEmmbedingUC embeddingGeneratorInterface) *spotUseCase {
	return &spotUseCase{
		repository:   repo,
		aiEmmbedding: aiEmmbedingUC,
	}
}

func (u *spotUseCase) GetAllSpots(ctx context.Context) ([]oapi.SpotResponse, error) {
	return u.repository.GetAllSpots(ctx)
}

func (u *spotUseCase) CreateSpot(ctx context.Context, spot *oapi.SpotResister) (oapi.SpotResponse, error) {
	var vector []float32
	// ベクトル化
	toVecText := spot.Name + "\n" + spot.Description + "\n" + spot.Address
	vector, err := u.aiEmmbedding.createEmbedding(ctx, toVecText)
	if err != nil {
		return oapi.SpotResponse{}, err
	}

	return u.repository.CreateSpot(ctx, spot, vector)
}

func (u *spotUseCase) GetSpotByID(ctx context.Context, spotId uuid.UUID) (oapi.SpotResponse, error) {
	return u.repository.GetSpotByID(ctx, spotId)
}

func (u *spotUseCase) UpdateSpotByID(ctx context.Context, spotId uuid.UUID, spot *oapi.SpotUpdate) (oapi.SpotResponse, error) {
	return u.repository.UpdateSpotByID(ctx, spotId, spot)
}
