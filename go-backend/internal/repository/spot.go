package repository

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

type SpotRepositoryInterface interface {
	GetAllSpots(ctx context.Context) ([]oapi.SpotResponse, error)
	CreateSpot(ctx context.Context, spot *oapi.SpotResister, vector []float32) (oapi.SpotResponse, error)
	GetSpotByID(ctx context.Context, spotId uuid.UUID) (oapi.SpotResponse, error)
	UpdateSpotByID(ctx context.Context, spotId uuid.UUID, spot *oapi.SpotUpdate) (oapi.SpotResponse, error)
	SearchSpots(ctx context.Context, query string) ([]oapi.SpotResponse, error)
	SearchSpotsByEmbedding(ctx context.Context, embedding []float32) ([]oapi.SpotResponse, error)
}

type spotRepositoryInmemory struct {
	mu      sync.RWMutex
	postsDB map[uuid.UUID]spotModel
}

type spotModel struct {
	Id          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Address     string    `db:"address"`
	CreatedAt   time.Time `db:"created_at"`
}

var _ SpotRepositoryInterface = (*spotRepositoryInmemory)(nil)

func NewSpotRepositoryInmemory() *spotRepositoryInmemory {
	return &spotRepositoryInmemory{
		// マップを初期化
		postsDB: make(map[uuid.UUID]spotModel),
	}
}

func (r *spotRepositoryInmemory) GetAllSpots(ctx context.Context) ([]oapi.SpotResponse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	AllSavedSpot := make([]oapi.SpotResponse, 0, len(r.postsDB))

	for _, spot := range r.postsDB {
		AllSavedSpot = append(AllSavedSpot, oapi.SpotResponse{
			Id:          spot.Id,
			Name:        spot.Name,
			Description: spot.Description,
			Address:     spot.Address,
			CreatedAt:   spot.CreatedAt.UTC(),
		})
	}

	return AllSavedSpot, nil
}

func (r *spotRepositoryInmemory) CreateSpot(ctx context.Context, spot *oapi.SpotResister, vector []float32) (oapi.SpotResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	newID := uuid.New()
	now := time.Now()

	newSpot := spotModel{
		Id:          newID,
		Name:        spot.Name,
		Description: spot.Description,
		Address:     spot.Address,
		CreatedAt:   now.UTC(),
	}

	r.postsDB[newID] = newSpot
	return oapi.SpotResponse{
		Id:          newSpot.Id,
		Name:        newSpot.Name,
		Description: newSpot.Description,
		Address:     newSpot.Address,
		CreatedAt:   newSpot.CreatedAt.UTC(),
	}, nil
}

func (r *spotRepositoryInmemory) GetSpotByID(ctx context.Context, spotId uuid.UUID) (oapi.SpotResponse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	spot, ok := r.postsDB[spotId]
	if !ok {
		return oapi.SpotResponse{}, fmt.Errorf("spot with ID %v not found", spotId)
	}

	return oapi.SpotResponse{
		Id:          spot.Id,
		Name:        spot.Name,
		Description: spot.Description,
		Address:     spot.Address,
		CreatedAt:   spot.CreatedAt.UTC(),
	}, nil
}

func (r *spotRepositoryInmemory) UpdateSpotByID(ctx context.Context, spotId uuid.UUID, spot *oapi.SpotUpdate) (oapi.SpotResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 更新対象のデータが存在するか確認
	existingSpot, ok := r.postsDB[spotId]
	if !ok {
		return oapi.SpotResponse{}, fmt.Errorf("spot with ID %v not found", spotId)
	}

	// データを更新
	if newName, err := spot.Name.Get(); err == nil {
		existingSpot.Name = newName
	}
	if newDescription, err := spot.Description.Get(); err == nil {
		existingSpot.Description = newDescription
	}
	if newAddress, err := spot.Address.Get(); err == nil {
		existingSpot.Address = newAddress
	}

	// マップに更新後のデータを保存
	r.postsDB[spotId] = existingSpot

	return oapi.SpotResponse{
		Id:          existingSpot.Id,
		Name:        existingSpot.Name,
		Description: existingSpot.Description,
		Address:     existingSpot.Address,
		CreatedAt:   existingSpot.CreatedAt.UTC(),
	}, nil
}

func (r *spotRepositoryInmemory) SearchSpots(ctx context.Context, query string) ([]oapi.SpotResponse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var results []oapi.SpotResponse
	for _, spot := range r.postsDB {
		if strings.Contains(strings.ToLower(spot.Name), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(spot.Description), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(spot.Address), strings.ToLower(query)) {
			results = append(results, oapi.SpotResponse{
				Id:          spot.Id,
				Name:        spot.Name,
				Description: spot.Description,
				Address:     spot.Address,
				CreatedAt:   spot.CreatedAt.UTC(),
			})
		}
	}
	return results, nil
}

func (r *spotRepositoryInmemory) SearchSpotsByEmbedding(ctx context.Context, embedding []float32) ([]oapi.SpotResponse, error) {
	// In-memory repository does not support vector search. Returning empty slice.
	return make([]oapi.SpotResponse, 0), nil
}
