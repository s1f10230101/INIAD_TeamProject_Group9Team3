package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository/sqlc"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// postgresReviewRepository は ReviewRepositoryInterface の PostgreSQL 実装です。
type postgresReviewRepository struct {
	db sqlc.DBTX
	q  *sqlc.Queries
}

// NewPostgresReviewRepository は新しい PostgresReviewRepository を生成します。
func NewPostgresReviewRepository(pool sqlc.DBTX) *postgresReviewRepository {
	return &postgresReviewRepository{
		db: pool,
		q:  sqlc.New(pool),
	}
}

// CreateReview は新しいレビューをDBに作成します。
func (r *postgresReviewRepository) CreateReview(ctx context.Context, spotId uuid.UUID, reviewInput *oapi.ReviewResister) (*oapi.ReviewResponse, error) {
	// TODO: 本来は認証情報から取得したユーザーIDを使います
	userID := uuid.New()

	params := sqlc.CreateReviewParams{
		ID:      uuid.New(),
		SpotID:  spotId,
		UserID:  userID,
		Rating:  int32(reviewInput.Rating),
		Comment: reviewInput.Comment,
		CreatedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	}

	createdSqlcReview, err := r.q.CreateReview(ctx, params)
	if err != nil {
		return nil, err
	}

	// sqlc.Review を api.Review に変換
	apiReview := &oapi.ReviewResponse{
		SpotId:    createdSqlcReview.SpotID,
		UserId:    createdSqlcReview.UserID,
		Rating:    int(createdSqlcReview.Rating),
		Comment:   createdSqlcReview.Comment,
		CreatedAt: createdSqlcReview.CreatedAt.Time.UTC(),
	}

	return apiReview, nil
}

// GetReviewsBySpotID は指定された観光施設のすべてのレビューをDBから取得します。
func (r *postgresReviewRepository) GetReviewsBySpotID(ctx context.Context, spotId uuid.UUID) ([]oapi.ReviewResponse, error) {
	sqlcReviews, err := r.q.GetReviewsBySpotID(ctx, spotId)
	if err != nil {
		return nil, err
	}

	// []sqlc.Review を []api.Review に変換
	apiReviews := make([]oapi.ReviewResponse, len(sqlcReviews))
	for i, sqlcReview := range sqlcReviews {
		apiReviews[i] = oapi.ReviewResponse{
			SpotId:    sqlcReview.SpotID,
			UserId:    sqlcReview.UserID,
			Rating:    int(sqlcReview.Rating),
			Comment:   sqlcReview.Comment,
			CreatedAt: sqlcReview.CreatedAt.Time.UTC(),
		}
	}

	return apiReviews, nil
}
