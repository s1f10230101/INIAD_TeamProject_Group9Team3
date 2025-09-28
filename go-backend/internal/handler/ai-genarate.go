package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// --- Python API Helper Structs ---

type EmbedRequest struct {
	Texts []string `json:"texts"`
}

type EmbedResponse struct {
	Status  string      `json:"status"`
	Vectors [][]float32 `json:"vectors"`
	Message string      `json:"message"`
}

type GenerateRequest struct {
	Question string `json:"question"`
	Context  string `json:"context"`
}

// GeneratePlan orchestrates the entire plan generation process.
// (POST /plans)
func (s *server) GeneratePlan(ctx context.Context, request oapi.GeneratePlanRequestObject) (oapi.GeneratePlanResponseObject, error) {
	prompt := request.Body.Prompt

	// --- 1. Vectorize the user's prompt ---
	embedURL := "http://python-backend:8000/embed"
	embedReqBody := EmbedRequest{Texts: []string{prompt}}
	jsonEmbedReq, err := json.Marshal(embedReqBody)
	if err != nil {
		return oapi.GeneratePlan500JSONResponse{Message: "Failed to create request body for embedding"}, err
	}

	embedResp, err := http.Post(embedURL, "application/json", bytes.NewBuffer(jsonEmbedReq))
	if err != nil {
		return oapi.GeneratePlan500JSONResponse{Message: fmt.Sprintf("Failed to call python embed service: %v", err)}, err
	}
	defer embedResp.Body.Close()

	if embedResp.StatusCode != http.StatusOK {
		return oapi.GeneratePlan500JSONResponse{Message: fmt.Sprintf("Python embed service returned error code: %d", embedResp.StatusCode)}, nil
	}

	var embedRespData EmbedResponse
	if err := json.NewDecoder(embedResp.Body).Decode(&embedRespData); err != nil {
		return oapi.GeneratePlan500JSONResponse{Message: "Failed to decode embedding response"}, err
	}

	if embedRespData.Status != "success" || len(embedRespData.Vectors) == 0 {
		return oapi.GeneratePlan500JSONResponse{Message: "Embedding failed or returned no vectors"}, nil
	}
	embedding := embedRespData.Vectors[0]

	// --- 2. Search for relevant spots in the database ---
	spots, err := s.spotRepo.SearchSpotsByVector(ctx, embedding)
	if err != nil {
		return oapi.GeneratePlan500JSONResponse{Message: fmt.Sprintf("Failed to search spots by vector: %v", err)}, err
	}

	// --- 3. Format the search results as context ---
	var contextBuilder strings.Builder
	if len(spots) == 0 {
		contextBuilder.WriteString("関連する観光スポットは見つかりませんでした。")
	} else {
		contextBuilder.WriteString("以下は関連する可能性のある観光スポットです。\n")
		for _, spot := range spots {
			contextBuilder.WriteString(fmt.Sprintf("- %s: %s\n", spot.Name, spot.Description))
		}
	}
	contextString := contextBuilder.String()

	// --- 4. Generate the travel plan using the context ---
	generateURL := "http://python-backend:8000/generate-plan"
	generateReqBody := GenerateRequest{
		Question: prompt,
		Context:  contextString,
	}
	jsonGenerateReq, err := json.Marshal(generateReqBody)
	if err != nil {
		return oapi.GeneratePlan500JSONResponse{Message: "Failed to create request body for generation"}, err
	}

	generateResp, err := http.Post(generateURL, "application/json", bytes.NewBuffer(jsonGenerateReq))
	if err != nil {
		return oapi.GeneratePlan500JSONResponse{Message: fmt.Sprintf("Failed to call python generate service: %v", err)}, err
	}

	if generateResp.StatusCode != http.StatusOK {
		defer generateResp.Body.Close()
		return oapi.GeneratePlan500JSONResponse{Message: fmt.Sprintf("Python generate service returned error code: %d", generateResp.StatusCode)}, nil
	}

	// --- 5. Stream the response back to the client ---
	return oapi.GeneratePlan200TexteventStreamResponse{
		Body:          generateResp.Body,
		ContentLength: generateResp.ContentLength,
	},
	nil
}