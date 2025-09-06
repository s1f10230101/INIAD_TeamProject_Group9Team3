package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/stretchr/testify/assert"
)

func TestPostReview(t *testing.T) {
	// --- setup ---
	// server_test.go のヘルパー関数を呼び出して、テスト用のルーターをセットアップ
	router := setupTestRouter()
	spotID := uuid.New()

	// --- execute ---
	// 1. レビューを投稿するためのリクエストを作成
	comment := "テストコメント: とても素晴らしい体験でした！"
	reviewInput := api.ReviewInput{
		Rating:  5,
		Comment: &comment,
	}
	body, _ := json.Marshal(reviewInput)
	reqPost := httptest.NewRequest(http.MethodPost, "/v1/spots/"+spotID.String()+"/reviews", strings.NewReader(string(body)))
	reqPost.Header.Set("Content-Type", "application/json")
	recPost := httptest.NewRecorder()

	// ルーターにリクエストを送信
	router.ServeHTTP(recPost, reqPost)

	// --- assert ---
	// レスポンスコードが201 Createdであることを確認
	assert.Equal(t, http.StatusCreated, recPost.Code)

	// レスポンスボディをパースして、内容が正しいか確認
	var createdReview api.Review
	err := json.Unmarshal(recPost.Body.Bytes(), &createdReview)
	assert.NoError(t, err)
	assert.Equal(t, reviewInput.Rating, createdReview.Rating)
	assert.Equal(t, *reviewInput.Comment, *createdReview.Comment)
	assert.Equal(t, spotID, createdReview.SpotId)
	assert.NotEmpty(t, createdReview.Id) // IDが生成されていることを確認

	// --- execute ---
	// 2. 投稿したレビューがGETで取得できるか確認
	reqGet := httptest.NewRequest(http.MethodGet, "/v1/spots/"+spotID.String()+"/reviews", nil)
	recGet := httptest.NewRecorder()
	router.ServeHTTP(recGet, reqGet)

	// --- assert ---
	assert.Equal(t, http.StatusOK, recGet.Code)
	var reviews []api.Review
	err = json.Unmarshal(recGet.Body.Bytes(), &reviews)
	assert.NoError(t, err)
	// 投稿したレビューが含まれていることを確認
	assert.Len(t, reviews, 1)
	assert.Equal(t, createdReview.Id, reviews[0].Id)
	assert.Equal(t, *createdReview.Comment, *reviews[0].Comment)
}
