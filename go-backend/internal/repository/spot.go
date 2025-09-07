package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
)

type SpotRepositoryInterface interface {
	GetAllSpots() ([]api.Spot, error)
	CreateSpot(spot *api.SpotInput) (api.Spot, error)
	GetSpotByID(spotId uuid.UUID) (api.Spot, error)
	UpdateSpotByID(spotId uuid.UUID, spot *api.SpotInput) (api.Spot, error)
}

type spotRepositoryInmemory struct {
	mu      sync.RWMutex
	postsDB map[uuid.UUID]api.Spot
}

func NewSpotRepositoryInmemory() *spotRepositoryInmemory {
	return &spotRepositoryInmemory{
		// マップを初期化
		postsDB: make(map[uuid.UUID]api.Spot),
	}
}

func (r *spotRepositoryInmemory) GetAllSpots() ([]api.Spot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	AllSavedSpot := make([]api.Spot, 0, len(r.postsDB))

	for _, spot := range r.postsDB {
		AllSavedSpot = append(AllSavedSpot, spot)
	}

	return AllSavedSpot, nil
}

func (r *spotRepositoryInmemory) CreateSpot(spot *api.SpotInput) (api.Spot, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	newID := uuid.New()
	now := time.Now()

	newSpot := api.Spot{
		Id:          newID,
		Name:        spot.Name,
		Description: &spot.Description,
		Address:     &spot.Address,
		CreatedAt:   now,
	}

	r.postsDB[newID] = newSpot
	return newSpot, nil
}

func (r *spotRepositoryInmemory) GetSpotByID(spotId uuid.UUID) (api.Spot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	//return r.postsDB[spotId], nil
	spot, ok := r.postsDB[spotId]
	if !ok {
		return api.Spot{}, fmt.Errorf("spot with ID %v not found", spotId)
	}

	return spot, nil
}

func (r *spotRepositoryInmemory) UpdateSpotByID(spotId uuid.UUID, spot *api.SpotInput) (api.Spot, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 更新対象のデータが存在するか確認
	existingSpot, ok := r.postsDB[spotId]
	if !ok {
		return api.Spot{}, fmt.Errorf("spot with ID %v not found", spotId)
	}

	// 既存のデータのフィールドを新しい情報で上書きする
	existingSpot.Name = spot.Name
	existingSpot.Description = &spot.Description
	existingSpot.Address = &spot.Address
	// IDとCreatedAtは変更しない

	// 更新したデータをマップに再保存する
	r.postsDB[spotId] = existingSpot

	// 更新後のデータを返す
	return existingSpot, nil
}
