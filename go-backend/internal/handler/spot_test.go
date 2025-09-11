package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

func TestPostAndGetSpot(t *testing.T) {
	// --- setup ---
	router := setupTestRouter()

	// --- execute POST ---
	// 1. 新しい観光施設を登録する
	spotInput := oapi.SpotResister{
		Name:        "テスト用観光地",
		Description: "これはハンドラテスト用の素晴らしい観光地です。",
		Address:     "東京都テスト区テスト1-1-1",
	}
	body, _ := json.Marshal(spotInput)
	reqPost := httptest.NewRequest(http.MethodPost, "/v1/spots", strings.NewReader(string(body)))
	reqPost.Header.Set("Content-Type", "application/json")
	recPost := httptest.NewRecorder()

	router.ServeHTTP(recPost, reqPost)

	if recPost.Code != http.StatusCreated {
		t.Fatalf("ステータスコードが%vではありません: got %v", http.StatusCreated, recPost.Code)
	}
	// レスポンスボディを検証
	var createdSpot oapi.SpotResponse
	err := json.Unmarshal(recPost.Body.Bytes(), &createdSpot)
	if err != nil {
		t.Fatalf("レスポンスのパースに失敗しました: %v", err)
	}

	if createdSpot.Name != spotInput.Name {
		t.Errorf("施設名が一致しません: got %v, want %v", createdSpot.Name, spotInput.Name)
	}
	if createdSpot.Description != spotInput.Description {
		t.Errorf("説明が一致しません: got %v, want %v", createdSpot.Description, spotInput.Description)
	}
	if createdSpot.Address != spotInput.Address {
		t.Errorf("住所が一致しません: got %v, want %v", createdSpot.Address, spotInput.Address)
	}
	if createdSpot.Id == uuid.Nil {
		t.Error("IDが生成されていません")
	}
	if createdSpot.CreatedAt.IsZero() {
		t.Error("CreatedAtが設定されていません")
	}

	// --- execute GET by ID ---
	// 2. IDを指定して施設を一件取得できるか確認
	spotID := createdSpot.Id.String()
	reqGet := httptest.NewRequest(http.MethodGet, "/v1/spots/"+spotID, nil)
	recGet := httptest.NewRecorder()

	router.ServeHTTP(recGet, reqGet)

	if recGet.Code != http.StatusOK {
		t.Fatalf("ステータスコードが200ではありません: got %v", recGet.Code)
	}

	// レスポンスボディを検証
	var fetchedSpot oapi.SpotResponse
	err = json.Unmarshal(recGet.Body.Bytes(), &fetchedSpot)
	if err != nil {
		t.Fatalf("レスポンスのパースに失敗しました: %v", err)
	}
	if createdSpot.Id != fetchedSpot.Id {
		t.Errorf("IDが一致しません: got %v, want %v", fetchedSpot.Id, createdSpot.Id)
	}
	if createdSpot.Name != fetchedSpot.Name {
		t.Errorf("施設名が一致しません: got %v, want %v", fetchedSpot.Name, createdSpot.Name)
	}
	if createdSpot.Description != fetchedSpot.Description {
		t.Errorf("説明が一致しません: got %v, want %v", fetchedSpot.Description, createdSpot.Description)
	}
}
