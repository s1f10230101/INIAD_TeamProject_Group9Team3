package usecase

import (
	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
)

// handler から呼び出されるビジネスロジックのインターフェース
type PostUseCaseInterface interface {
	GetAllSpots() ([]api.Spot, error)
	CreateSpot(spot *api.SpotInput) error
	GetSpotByID(spotId uuid.UUID) (api.Spot, error)
	UpdateSpotByID(spotId uuid.UUID, spot *api.SpotInput) (api.Spot, error)
}

type postUseCase struct {
	repository repository.PostRepositoryInterface
}

func NewPostUseCase(repo repository.PostRepositoryInterface) *postUseCase {
	return &postUseCase{
		repository: repo,
	}
}

func (u *postUseCase) GetAllSpots() ([]api.Spot, error) {
	return u.repository.GetAllSpots()
}

func (u *postUseCase) CreateSpot(spot *api.SpotInput) error {
	return u.repository.CreateSpot(spot)
}

func (u *postUseCase) GetSpotByID(spotId uuid.UUID) (api.Spot, error) {
	return u.repository.GetSpotByID(spotId)
}

func (u *postUseCase) UpdateSpotByID(spotId uuid.UUID, spot *api.SpotInput) (api.Spot, error) {
	return u.repository.UpdateSpotByID(spotId, spot)
}
