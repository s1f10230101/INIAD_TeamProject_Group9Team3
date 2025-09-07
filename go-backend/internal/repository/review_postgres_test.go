package repository

import (
	"context"
	"testing"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository/sqlc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostgresReviewRepository_CreateAndGetReviews(t *testing.T) {
	// --- Setup: トランザクションを開始 ---
	tx, err := testPool.Begin(context.Background())
	require.NoError(t, err)
	defer tx.Rollback(context.Background())

	// --- Setup: 同じトランザクションでリポジトリを初期化 ---
	q := sqlc.New(tx)
	spotRepo := &postgresSpotRepository{q: q}
	reviewRepo := &PostgresReviewRepository{db: q}

	// --- Setup: レビュー対象のSpotを作成 ---
	spotInput := &api.SpotInput{
		Name:        "レビューテスト用の観光地",
		Description: "レビューのインテグレーションテスト",
		Address:     "東京都",
	}
	createdSpot, err := spotRepo.CreateSpot(context.Background(), spotInput)
	require.NoError(t, err)

	// --- Test: CreateReview ---
	comment := "素晴らしい場所でした！"
	reviewInput := &api.ReviewInput{
		Rating:  5,
		Comment: &comment,
	}

	createdReview, err := reviewRepo.CreateReview(context.Background(), createdSpot.Id, reviewInput)
	require.NoError(t, err)

	// --- Assert: CreateReviewの結果を検証 ---
	assert.NotEmpty(t, createdReview.Id)
	assert.Equal(t, createdSpot.Id, createdReview.SpotId)
	assert.Equal(t, reviewInput.Rating, createdReview.Rating)
	assert.Equal(t, *reviewInput.Comment, *createdReview.Comment)
	assert.NotEmpty(t, createdReview.UserId)
	assert.NotZero(t, createdReview.CreatedAt)

	// --- Test: GetReviewsBySpotID ---
	reviews, err := reviewRepo.GetReviewsBySpotID(context.Background(), createdSpot.Id)
	require.NoError(t, err)

	// --- Assert: GetReviewsBySpotIDの結果を検証 ---
	assert.Len(t, reviews, 1)
	assert.Equal(t, createdReview.Id, reviews[0].Id)
	assert.Equal(t, *createdReview.Comment, *reviews[0].Comment)
}
