package repository

import (
	"sync"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
)

type PostRepositoryInterface interface {
	GetAllSpots() ([]api.Spot, error)
	CreateSpot(spot *api.SpotInput) error
	GetSpotByID(spotId uuid.UUID) (api.Spot, error)
	UpdateSpotByID(spotId uuid.UUID, spot *api.SpotInput) (api.Spot, error)
}

type postRepositoryInmemory struct {
	mu      sync.RWMutex
	postsDB map[uuid.UUID]api.Spot
}

func NewPostRepositoryInmemory() *postRepositoryInmemory {
	return &postRepositoryInmemory{}
}

func (r *postRepositoryInmemory) GetAllSpots() ([]api.Spot, error)

func (r *postRepositoryInmemory) CreateSpot(spot *api.SpotInput) error

func (r *postRepositoryInmemory) GetSpotByID(spotId uuid.UUID) (api.Spot, error)

func (r *postRepositoryInmemory) UpdateSpotByID(spotId uuid.UUID, spot *api.SpotInput) (api.Spot, error)
