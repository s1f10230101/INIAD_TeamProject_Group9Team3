package repository

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// ReviewRepositoryInterface はレビューデータの永続化を担うインターフェースです。
type ReviewRepositoryInterface interface {
	CreateReview(ctx context.Context, spotId uuid.UUID, reviewInput *oapi.ReviewResister) (*oapi.ReviewResponse, error)
	GetReviewsBySpotID(ctx context.Context, spotId uuid.UUID) ([]oapi.ReviewResponse, error)
}

type reviewModel struct {
	Id        uuid.UUID `db:"id"`
	SpotId    uuid.UUID `db:"spot_id"`
	UserId    uuid.UUID `db:"user_id"`
	Rating    int       `db:"rating"`
	Comment   string    `db:"comment"`
	CreatedAt time.Time `db:"created_at"`
}

type reviewRepositoryInmemory struct {
	mu        sync.RWMutex
	reviewsDB map[uuid.UUID]reviewModel
}

// NewReviewRepositoryInmemory はインメモリのレビューリポジトリを生成します。
func NewReviewRepositoryInmemory() ReviewRepositoryInterface {
	return &reviewRepositoryInmemory{
		reviewsDB: make(map[uuid.UUID]reviewModel),
	}
}

// GetReviewsBySpotID は指定された観光施設のすべてのレビューを取得します。
func (r *reviewRepositoryInmemory) GetReviewsBySpotID(ctx context.Context, spotId uuid.UUID) ([]oapi.ReviewResponse, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var reviews []oapi.ReviewResponse
	// マップをループしてspotIdが一致するものを探す
	for _, review := range r.reviewsDB {
		if review.SpotId == spotId {
			reviews = append(reviews, oapi.ReviewResponse{
				SpotId:  review.SpotId,
				UserId:  review.UserId,
				Rating:  review.Rating,
				Comment: review.Comment,
			})
		}
	}
	return reviews, nil
}

// CreateReview は新しいレビューを作成します。
func (r *reviewRepositoryInmemory) CreateReview(ctx context.Context, spotId uuid.UUID, reviewInput *oapi.ReviewResister) (*oapi.ReviewResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	newReview := reviewModel{
		Id:        uuid.New(),
		SpotId:    spotId,
		UserId:    uuid.New(), // 補足: 本来は認証情報から取得したユーザーIDを使います
		Rating:    reviewInput.Rating,
		Comment:   reviewInput.Comment,
		CreatedAt: time.Now(),
	}

	r.reviewsDB[newReview.Id] = newReview
	return &oapi.ReviewResponse{
		SpotId:  newReview.SpotId,
		UserId:  newReview.UserId,
		Rating:  newReview.Rating,
		Comment: newReview.Comment,
	}, nil
}
