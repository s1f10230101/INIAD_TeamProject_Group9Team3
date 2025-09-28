package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// --- Python API Helper Structs ---
// Note: These are duplicated from the handler. In a larger app, they could be moved to a shared package.
type EmbedRequest struct {
	Texts []string `json:"texts"`
}

type EmbedResponse struct {
	Status  string      `json:"status"`
	Vectors [][]float32 `json:"vectors"`
	Message string      `json:"message"`
}

// SpotUseCaseInterface defines the business logic for spots.
type SpotUseCaseInterface interface {
	GetAllSpots(ctx context.Context) ([]oapi.SpotResponse, error)
	CreateSpot(ctx context.Context, spot *oapi.SpotResister) (oapi.SpotResponse, error)
	GetSpotByID(ctx context.Context, spotId uuid.UUID) (oapi.SpotResponse, error)
	UpdateSpotByID(ctx context.Context, spotId uuid.UUID, spot *oapi.SpotUpdate) (oapi.SpotResponse, error)
}

type spotUseCase struct {
	repository repository.SpotRepositoryInterface
}

var _ SpotUseCaseInterface = (*spotUseCase)(nil)

func NewSpotUseCase(repo repository.SpotRepositoryInterface) *spotUseCase {
	return &spotUseCase{
		repository: repo,
	}
}

func (u *spotUseCase) GetAllSpots(ctx context.Context) ([]oapi.SpotResponse, error) {
	return u.repository.GetAllSpots(ctx)
}

func (u *spotUseCase) CreateSpot(ctx context.Context, spot *oapi.SpotResister) (oapi.SpotResponse, error) {
	// 1. Create the spot in the database first
	createdSpot, err := u.repository.CreateSpot(ctx, spot)
	if err != nil {
		return oapi.SpotResponse{}, err
	}

	// 2. Asynchronously generate and update the embedding
	go u.generateAndUpdateEmbedding(context.Background(), createdSpot.Id, createdSpot.Name, createdSpot.Description)

	return createdSpot, nil
}

func (u *spotUseCase) GetSpotByID(ctx context.Context, spotId uuid.UUID) (oapi.SpotResponse, error) {
	return u.repository.GetSpotByID(ctx, spotId)
}

func (u *spotUseCase) UpdateSpotByID(ctx context.Context, spotId uuid.UUID, spot *oapi.SpotUpdate) (oapi.SpotResponse, error) {
	// 1. Update the spot in the database first
	updatedSpot, err := u.repository.UpdateSpotByID(ctx, spotId, spot)
	if err != nil {
		return oapi.SpotResponse{}, err
	}

	// 2. Asynchronously generate and update the embedding
	go u.generateAndUpdateEmbedding(context.Background(), updatedSpot.Id, updatedSpot.Name, updatedSpot.Description)

	return updatedSpot, nil
}

// generateAndUpdateEmbedding is a helper function to handle the embedding process.
// It runs in a separate goroutine so it doesn't block the user's request.
func (u *spotUseCase) generateAndUpdateEmbedding(ctx context.Context, spotId uuid.UUID, name, description string) {
	log.Printf("Starting embedding generation for spot ID: %s", spotId)

	// 1. Prepare text for embedding
	textToEmbed := fmt.Sprintf("観光地名: %s\n説明: %s", name, description)

	// 2. Call Python's /embed service
	embedURL := "http://python-backend:8000/embed"
	embedReqBody := EmbedRequest{Texts: []string{textToEmbed}}
	jsonEmbedReq, err := json.Marshal(embedReqBody)
	if err != nil {
		log.Printf("ERROR: Failed to create embedding request body for spot %s: %v", spotId, err)
		return
	}

	embedResp, err := http.Post(embedURL, "application/json", bytes.NewBuffer(jsonEmbedReq))
	if err != nil {
		log.Printf("ERROR: Failed to call python embed service for spot %s: %v", spotId, err)
		return
	}
	defer embedResp.Body.Close()

	if embedResp.StatusCode != http.StatusOK {
		log.Printf("ERROR: Python embed service returned error code %d for spot %s", embedResp.StatusCode, spotId)
		return
	}

	var embedRespData EmbedResponse
	if err := json.NewDecoder(embedResp.Body).Decode(&embedRespData); err != nil {
		log.Printf("ERROR: Failed to decode embedding response for spot %s: %v", spotId, err)
		return
	}

	if embedRespData.Status != "success" || len(embedRespData.Vectors) == 0 {
		log.Printf("ERROR: Embedding failed or returned no vectors for spot %s. Message: %s", spotId, embedRespData.Message)
		return
	}
	embedding := embedRespData.Vectors[0]

	// 3. Update the spot in the database with the new embedding
	err = u.repository.UpdateSpotEmbedding(ctx, spotId, embedding)
	if err != nil {
		log.Printf("ERROR: Failed to update spot embedding for spot %s: %v", spotId, err)
		return
	}

	log.Printf("Successfully updated embedding for spot ID: %s", spotId)
}

func NewPostUseCase(repo repository.SpotRepositoryInterface) SpotUseCaseInterface {
	return &spotUseCase{
		repository: repo,
	}
}
