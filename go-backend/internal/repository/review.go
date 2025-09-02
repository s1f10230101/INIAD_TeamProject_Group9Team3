package repository

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
)

// ReviewRepositoryInterface はレビューデータの永続化を担うインターフェースです。
type ReviewRepositoryInterface interface {
	GetReviewsBySpotID(spotId uuid.UUID) ([]api.Review, error)
	CreateReview(spotId uuid.UUID, review *api.ReviewInput) (*api.Review, error)
}

type reviewRepositoryInmemory struct {
	mu        sync.RWMutex
	reviewsDB map[uuid.UUID]api.Review
}

// NewReviewRepositoryInmemory はインメモリのレビューリポジトリを生成します。
func NewReviewRepositoryInmemory() ReviewRepositoryInterface {
	return &reviewRepositoryInmemory{
		reviewsDB: make(map[uuid.UUID]api.Review),
	}
}

// GetReviewsBySpotID は指定された観光施設のすべてのレビューを取得します。
func (r *reviewRepositoryInmemory) GetReviewsBySpotID(spotId uuid.UUID) ([]api.Review, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var reviews []api.Review
	// マップをループしてspotIdが一致するものを探す
	for _, review := range r.reviewsDB {
		if review.SpotId == spotId {
			reviews = append(reviews, review)
		}
	}
	return reviews, nil
}

// CreateReview は新しいレビューを作成します。
func (r *reviewRepositoryInmemory) CreateReview(spotId uuid.UUID, reviewInput *api.ReviewInput) (*api.Review, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	newReview := api.Review{
		Id:        uuid.New(),
		SpotId:    spotId,
		UserId:    uuid.New(), // 補足: 本来は認証情報から取得したユーザーIDを使います
		Rating:    reviewInput.Rating,
		Comment:   reviewInput.Comment,
		CreatedAt: time.Now(),
	}

	r.reviewsDB[newReview.Id] = newReview
	return &newReview, nil
}
