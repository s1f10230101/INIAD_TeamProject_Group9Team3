package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/stretchr/testify/assert"
)

func TestPostAndGetSpot(t *testing.T) {
	// --- setup ---
	router := setupTestRouter()

	// --- execute POST ---
	// 1. 新しい観光施設を登録する
	spotInput := api.SpotInput{
		Name:        "テスト用観光地",
		Description: "これはハンドラテスト用の素晴らしい観光地です。",
		Address:     "東京都テスト区テスト1-1-1",
	}
	body, _ := json.Marshal(spotInput)
	reqPost := httptest.NewRequest(http.MethodPost, "/v1/spots", strings.NewReader(string(body)))
	reqPost.Header.Set("Content-Type", "application/json")
	recPost := httptest.NewRecorder()

	router.ServeHTTP(recPost, reqPost)

	// --- assert POST ---
	assert.Equal(t, http.StatusCreated, recPost.Code)

	// レスポンスボディを検証
	var createdSpot api.Spot
	err := json.Unmarshal(recPost.Body.Bytes(), &createdSpot)
	assert.NoError(t, err)

	assert.Equal(t, spotInput.Name, createdSpot.Name)
	assert.Equal(t, spotInput.Description, *createdSpot.Description)
	assert.Equal(t, spotInput.Address, *createdSpot.Address)
	assert.NotEmpty(t, createdSpot.Id, "ID should be generated")
	assert.False(t, createdSpot.CreatedAt.IsZero(), "CreatedAt should be set")

	// --- execute GET by ID ---
	// 2. IDを指定して施設を一件取得できるか確認
	spotID := createdSpot.Id.String()
	reqGet := httptest.NewRequest(http.MethodGet, "/v1/spots/"+spotID, nil)
	recGet := httptest.NewRecorder()

	router.ServeHTTP(recGet, reqGet)

	// --- assert GET by ID ---
	assert.Equal(t, http.StatusOK, recGet.Code)

	var fetchedSpot api.Spot
	err = json.Unmarshal(recGet.Body.Bytes(), &fetchedSpot)
	assert.NoError(t, err)
	assert.Equal(t, createdSpot.Id, fetchedSpot.Id)
	assert.Equal(t, createdSpot.Name, fetchedSpot.Name)
	assert.Equal(t, *createdSpot.Description, *fetchedSpot.Description)
}
