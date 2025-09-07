package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository/sqlc"
)

type postgresSpotRepository struct {
	db *pgxpool.Pool
	q  *sqlc.Queries
}

// NewPostgresPostRepository は新しいpostgresPostRepositoryのインスタンスを生成
func NewPostgresPostRepository(db *pgxpool.Pool) SpotRepositoryInterface {
	return &postgresSpotRepository{
		db: db,
		q:  sqlc.New(db),
	}
}

func (r *postgresSpotRepository) GetAllSpots(ctx context.Context) ([]api.Spot, error) {
	rows, err := r.q.ListSpots(ctx)
	if err != nil {
		return nil, err
	}

	spots := make([]api.Spot, len(rows))
	for i, row := range rows {
		spots[i] = api.Spot{
			Id:          row.ID,
			Name:        row.Name,
			Description: &row.Description,
			Address:     &row.Address,
			CreatedAt:   row.CreatedAt.Time,
		}
	}
	return spots, nil
}

func (r *postgresSpotRepository) CreateSpot(ctx context.Context, spot *api.SpotInput) (api.Spot, error) {
	newID := uuid.New()
	params := sqlc.CreateSpotParams{
		ID:          newID,
		Name:        spot.Name,
		Description: spot.Description,
		Address:     spot.Address,
	}
	created, err := r.q.CreateSpot(ctx, params)
	if err != nil {
		return api.Spot{}, err
	}

	return api.Spot{
		Id:          created.ID,
		Name:        created.Name,
		Description: &created.Description,
		Address:     &created.Address,
		CreatedAt:   created.CreatedAt.Time,
	}, nil
}

func (r *postgresSpotRepository) GetSpotByID(ctx context.Context, spotId uuid.UUID) (api.Spot, error) {
	row, err := r.q.GetSpot(ctx, spotId)
	if err != nil {
		return api.Spot{}, err
	}
	return api.Spot{
		Id:          row.ID,
		Name:        row.Name,
		Description: &row.Description,
		Address:     &row.Address,
		CreatedAt:   row.CreatedAt.Time,
	}, nil
}

func (r *postgresSpotRepository) UpdateSpotByID(ctx context.Context, spotId uuid.UUID, spot *api.SpotInput) (api.Spot, error) {
	params := sqlc.UpdateSpotParams{
		ID:          spotId,
		Name:        spot.Name,
		Description: spot.Description,
		Address:     spot.Address,
	}
	updated, err := r.q.UpdateSpot(ctx, params)
	if err != nil {
		return api.Spot{}, err
	}
	return api.Spot{
		Id:          updated.ID,
		Name:        updated.Name,
		Description: &updated.Description,
		Address:     &updated.Address,
		CreatedAt:   updated.CreatedAt.Time,
	}, nil
}


