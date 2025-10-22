package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// ReviewUseCaseInterface はレビューに関するビジネスロジックのインターフェースです。
type ReviewUseCaseInterface interface {
	GetReviewsBySpotID(ctx context.Context, spotId uuid.UUID) ([]oapi.ReviewResponse, error)
	CreateReview(ctx context.Context, spotId uuid.UUID, review *oapi.ReviewResister) (*oapi.ReviewResponse, error)
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

func (u *reviewUseCase) GetReviewsBySpotID(ctx context.Context, spotId uuid.UUID) ([]oapi.ReviewResponse, error) {
	return u.reviewRepo.GetReviewsBySpotID(ctx, spotId)
}

func (u *reviewUseCase) CreateReview(ctx context.Context, spotId uuid.UUID, review *oapi.ReviewResister) (*oapi.ReviewResponse, error) {
	if review.Rating < 1 || review.Rating > 6 {
		return nil, fmt.Errorf("review 1~5")
	}
	return u.reviewRepo.CreateReview(ctx, spotId, review)
}
