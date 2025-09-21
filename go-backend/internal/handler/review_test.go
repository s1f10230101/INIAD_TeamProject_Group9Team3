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

func TestPostReview(t *testing.T) {
	// --- setup ---
	// server_test.go のヘルパー関数を呼び出して、テスト用のルーターをセットアップ
	router := setupTestRouter()
	spotID := uuid.New()

	// --- execute ---
	// 1. レビューを投稿するためのリクエストを作成
	comment := "テストコメント: とても素晴らしい体験でした！"
	reviewInput := oapi.ReviewResister{
		Rating:  59,
		Comment: comment,
	}
	body, _ := json.Marshal(reviewInput)
	reqPost := httptest.NewRequest(http.MethodPost, "/v1/spots/"+spotID.String()+"/reviews", strings.NewReader(string(body)))
	reqPost.Header.Set("Content-Type", "application/json")
	recPost := httptest.NewRecorder()

	// ルーターにリクエストを送信
	router.ServeHTTP(recPost, reqPost)

	// レスポンスコードが201 Createdであることを確認
	if recPost.Code != http.StatusCreated {
		t.Fatalf("期待されるステータスコード %d ではありません: got %d", http.StatusCreated, recPost.Code)
	}

	// レスポンスボディをパースして、内容が正しいか確認
	var createdReview oapi.ReviewResponse
	err := json.Unmarshal(recPost.Body.Bytes(), &createdReview)
	if err != nil {
		t.Fatalf("レスポンスのパースに失敗しました: %v", err)
	}
	if createdReview.Comment != comment {
		t.Errorf("Expected comment %q, but got %q", comment, createdReview.Comment)
	}
	if createdReview.Rating != reviewInput.Rating {
		t.Errorf("Expected rating %d, but got %d", reviewInput.Rating, createdReview.Rating)
	}
	if createdReview.Comment != reviewInput.Comment {
		t.Errorf("Expected comment %q, but got %q", reviewInput.Comment, createdReview.Comment)
	}
	if createdReview.SpotId != spotID {
		t.Errorf("Expected SpotId %v, but got %v", spotID, createdReview.SpotId)
	}
	if createdReview.SpotId == uuid.Nil {
		t.Error("IDが生成されていません")
	}
	// --- execute ---
	// 2. 投稿したレビューがGETで取得できるか確認
	reqGet := httptest.NewRequest(http.MethodGet, "/v1/spots/"+spotID.String()+"/reviews", nil)
	recGet := httptest.NewRecorder()
	router.ServeHTTP(recGet, reqGet)

	if recGet.Code != http.StatusOK {
		t.Fatalf("期待されるステータスコード %d ではありません: got %d", http.StatusOK, recGet.Code)
	}
	var reviews []oapi.ReviewResponse
	err = json.Unmarshal(recGet.Body.Bytes(), &reviews)
	if err != nil {
		t.Fatalf("レスポンスのパースに失敗しました: %v", err)
	}
	// 投稿したレビューが含まれていることを確認
	if reviews[0].SpotId != createdReview.SpotId {
		t.Errorf("レビューIDが一致しません: got %v, want %v", reviews[0].SpotId, createdReview.SpotId)
	}
	if reviews[0].Rating != createdReview.Rating {
		t.Errorf("評価が一致しません: got %v, want %v", reviews[0].Rating, createdReview.Rating)
	}
	if createdReview.Comment != reviews[0].Comment {
		t.Errorf("コメントが一致しません: got %v, want %v", reviews[0].Comment, createdReview.Comment)
	}
}
