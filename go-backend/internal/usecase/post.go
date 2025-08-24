package usecase

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
)

type PostUseCaseInterface interface {
	GetAllSpots() ([]api.Spot, error)
	CreateSpot(spot *api.SpotInput) error
	GetSpotByID(spotId openapi_types.UUID) (api.Spot, error)
	UpdateSpotByID(spotId openapi_types.UUID, spot *api.SpotInput) (api.Spot, error)
}

type postUseCase struct {
	repository repository.PostRepositoryInterface
}

func NewPostUseCase(repo repository.PostRepositoryInterface) PostUseCaseInterface {
	return &postUseCase{
		repository: repo,
	}
}

func (u *postUseCase) GetAllSpots() ([]api.Spot, error)

func (u *postUseCase) CreateSpot(spot *api.SpotInput) error

func (u *postUseCase) GetSpotByID(spotId openapi_types.UUID) (api.Spot, error)

func (u *postUseCase) UpdateSpotByID(spotId openapi_types.UUID, spot *api.SpotInput) (api.Spot, error)
