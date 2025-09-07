package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository/sqlc"
)

// PostgresReviewRepository は ReviewRepositoryInterface の PostgreSQL 実装です。
type PostgresReviewRepository struct {
	db *sqlc.Queries
}

// NewPostgresReviewRepository は新しい PostgresReviewRepository を生成します。
func NewPostgresReviewRepository(pool *pgxpool.Pool) ReviewRepositoryInterface {
	return &PostgresReviewRepository{
		db: sqlc.New(pool),
	}
}

// CreateReview は新しいレビューをDBに作成します。
func (r *PostgresReviewRepository) CreateReview(spotId uuid.UUID, reviewInput *api.ReviewInput) (*api.Review, error) {
	ctx := context.Background()

	// TODO: 本来は認証情報から取得したユーザーIDを使います
	userID := uuid.New()

	params := sqlc.CreateReviewParams{
		ID:      uuid.New(),
		SpotID:  spotId,
		UserID:  userID,
		Rating:  int32(reviewInput.Rating),
		Comment: *reviewInput.Comment,
		CreatedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	}

	createdSqlcReview, err := r.db.CreateReview(ctx, params)
	if err != nil {
		return nil, err
	}

	// sqlc.Review を api.Review に変換
	apiReview := &api.Review{
		Id:        createdSqlcReview.ID,
		SpotId:    createdSqlcReview.SpotID,
		UserId:    createdSqlcReview.UserID,
		Rating:    int32(createdSqlcReview.Rating),
		Comment:   &createdSqlcReview.Comment,
		CreatedAt: createdSqlcReview.CreatedAt.Time,
	}

	return apiReview, nil
}

// GetReviewsBySpotID は指定された観光施設のすべてのレビューをDBから取得します。
func (r *PostgresReviewRepository) GetReviewsBySpotID(spotId uuid.UUID) ([]api.Review, error) {
	ctx := context.Background()

	sqlcReviews, err := r.db.GetReviewsBySpotID(ctx, spotId)
	if err != nil {
		return nil, err
	}

	// []sqlc.Review を []api.Review に変換
	apiReviews := make([]api.Review, len(sqlcReviews))
	for i, sqlcReview := range sqlcReviews {
		apiReviews[i] = api.Review{
			Id:        sqlcReview.ID,
			SpotId:    sqlcReview.SpotID,
			UserId:    sqlcReview.UserID,
			Rating:    int32(sqlcReview.Rating),
			Comment:   &sqlcReview.Comment,
			CreatedAt: sqlcReview.CreatedAt.Time,
		}
	}

	return apiReviews, nil
}


