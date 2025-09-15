//go:build integration

package repository_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

func TestPostgresReviewRepository_CreateAndGetReviews(t *testing.T) {
	// --- Setup: トランザクションを開始 ---
	ctx := context.Background()
	tx, err := testPool.Begin(ctx)
	if err != nil {
		t.Fatalf("トランザクションの開始に失敗: %v", err)
	}
	defer tx.Rollback(ctx)

	// --- Setup: 同じトランザクションでリポジトリを初期化 ---
	spotRepo := repository.NewPostgresPostRepository(tx)
	reviewRepo := repository.NewPostgresReviewRepository(tx)

	// --- Setup: レビュー対象のSpotを作成 ---
	spotInput := &oapi.SpotResister{
		Name:        "レビューテスト用の観光地",
		Description: "レビューのインテグレーションテスト",
		Address:     "東京都",
	}
	createdSpot, err := spotRepo.CreateSpot(context.Background(), spotInput, nil)
	if err != nil {
		t.Fatalf("スポットの作成に失敗: %v", err)
	}

	// --- Test: CreateReview ---
	comment := "素晴らしい場所でした！"
	reviewInput := &oapi.ReviewResister{
		Rating:  5,
		Comment: comment,
	}

	createdReview, err := reviewRepo.CreateReview(context.Background(), createdSpot.Id, reviewInput)
	if err != nil {
		t.Fatalf("レビューの作成に失敗: %v", err)
	}

	// --- Assert: CreateReviewの結果を検証 ---
	if createdReview.SpotId != createdSpot.Id {
		t.Errorf("SpotIdが一致しません: got %v, want %v", createdReview.SpotId, createdSpot.Id)
	}
	if createdReview.Rating != reviewInput.Rating {
		t.Errorf("Ratingが一致しません: got %v, want %v", createdReview.Rating, reviewInput.Rating)
	}
	if createdReview.Comment != reviewInput.Comment {
		t.Errorf("Commentが一致しません: got %v, want %v", createdReview.Comment, reviewInput.Comment)
	}
	if createdReview.UserId == uuid.Nil {
		t.Error("UserIdが設定されていません")
	}
	if createdReview.CreatedAt.IsZero() {
		t.Error("CreatedAtが設定されていません")
	}
	if createdReview.SpotId == uuid.Nil {
		t.Error("IDが設定されていません")
	}

	// --- Test: GetReviewsBySpotID ---
	reviews, err := reviewRepo.GetReviewsBySpotID(context.Background(), createdSpot.Id)
	if err != nil {
		t.Fatalf("レビューの取得に失敗: %v", err)
	}

	// --- Assert: GetReviewsBySpotIDの結果を検証 ---
	if reviews[0].SpotId != createdReview.SpotId {
		t.Errorf("レビューIDが一致しません: got %v, want %v", reviews[0].SpotId, createdReview.SpotId)
	}
	if reviews[0].Rating != createdReview.Rating {
		t.Errorf("評価が一致しません: got %v, want %v", reviews[0].Rating, createdReview.Rating)
	}
	if reviews[0].Comment != createdReview.Comment {
		t.Errorf("コメントが一致しません: got %v, want %v", reviews[0].Comment, createdReview.Comment)
	}
}
