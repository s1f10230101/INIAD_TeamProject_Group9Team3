package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository/sqlc"
)

type postgresSpostRepository struct {
	db *pgxpool.Pool
	q  *sqlc.Queries
}

// NewPostgresPostRepository は新しいpostgresPostRepositoryのインスタンスを生成
func NewPostgresPostRepository(db *pgxpool.Pool) PostRepositoryInterface {
	return &postgresSpostRepository{
		db: db,
		q:  sqlc.New(db),
	}
}

func (r *postgresSpostRepository) GetAllSpots() ([]api.Spot, error) {
	rows, err := r.q.ListSpots(context.Background())
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

func (r *postgresSpostRepository) CreateSpot(spot *api.SpotInput) (api.Spot, error) {
	newID := uuid.New()
	params := sqlc.CreateSpotParams{
		ID:          newID,
		Name:        spot.Name,
		Description: spot.Description,
		Address:     spot.Address,
	}
	created, err := r.q.CreateSpot(context.Background(), params)
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

func (r *postgresSpostRepository) GetSpotByID(spotId uuid.UUID) (api.Spot, error) {
	row, err := r.q.GetSpot(context.Background(), spotId)
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

func (r *postgresSpostRepository) UpdateSpotByID(spotId uuid.UUID, spot *api.SpotInput) (api.Spot, error) {
	params := sqlc.UpdateSpotParams{
		ID:          spotId,
		Name:        spot.Name,
		Description: spot.Description,
		Address:     spot.Address,
	}
	updated, err := r.q.UpdateSpot(context.Background(), params)
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

// NewPostgresPostRepositoryForTest はテスト用にトランザクションを受け取るコンストラクタです。
// この関数は repository パッケージのテストでのみ使用されることを想定しています。
func NewPostgresPostRepositoryForTest(tx sqlc.DBTX) PostRepositoryInterface {
	return &postgresSpostRepository{
		db: nil,
		q:  sqlc.New(tx),
	}
}
