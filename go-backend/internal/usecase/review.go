package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
)

// ReviewUseCaseInterface はレビューに関するビジネスロジックのインターフェースです。
type ReviewUseCaseInterface interface {
	GetReviewsBySpotID(ctx context.Context, spotId uuid.UUID) ([]api.Review, error)
	CreateReview(ctx context.Context, spotId uuid.UUID, review *api.ReviewInput) (*api.Review, error)
}

type reviewUseCase struct {
	reviewRepo repository.ReviewRepositoryInterface
}

// NewReviewUseCase は新しいレビューユースケースのインスタンスを生成します。
func NewReviewUseCase(repo repository.ReviewRepositoryInterface) ReviewUseCaseInterface {
	return &reviewUseCase{
		reviewRepo: repo,
	}
}

func (u *reviewUseCase) GetReviewsBySpotID(ctx context.Context, spotId uuid.UUID) ([]api.Review, error) {
	return u.reviewRepo.GetReviewsBySpotID(ctx, spotId)
}

func (u *reviewUseCase) CreateReview(ctx context.Context, spotId uuid.UUID, review *api.ReviewInput) (*api.Review, error) {
	return u.reviewRepo.CreateReview(ctx, spotId, review)
}
