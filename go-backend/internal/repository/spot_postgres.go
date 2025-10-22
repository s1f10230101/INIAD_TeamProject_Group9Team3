package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository/sqlc"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

type postgresSpotRepository struct {
	db sqlc.DBTX
	q  *sqlc.Queries
}

var _ SpotRepositoryInterface = (*postgresSpotRepository)(nil)

// NewPostgresPostRepository は新しいpostgresPostRepositoryのインスタンスを生成
func NewPostgresPostRepository(pool sqlc.DBTX) *postgresSpotRepository {
	return &postgresSpotRepository{
		db: pool,
		q:  sqlc.New(pool),
	}
}

func (r *postgresSpotRepository) GetAllSpots(ctx context.Context) ([]oapi.SpotResponse, error) {
	rows, err := r.q.ListSpots(ctx)
	if err != nil {
		return nil, err
	}

	spots := make([]oapi.SpotResponse, len(rows))
	for i, row := range rows {
		spots[i] = oapi.SpotResponse{
			Id:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			Address:     row.Address,
			CreatedAt:   row.CreatedAt.Time.UTC(),
		}
	}
	return spots, nil
}

func (r *postgresSpotRepository) CreateSpot(ctx context.Context, spot *oapi.SpotResister, vector []float32) (oapi.SpotResponse, error) {
	newID := uuid.New()

	// Embeddingを生成
	pgvecEmbedding := pgvector.NewVector(vector)

	params := sqlc.CreateSpotParams{
		ID:              newID,
		Name:            spot.Name,
		Description:     spot.Description,
		Address:         spot.Address,
		EmbeddingOpenai: &pgvecEmbedding,
	}
	createdSpot, err := r.q.CreateSpot(ctx, params)
	if err != nil {
		return oapi.SpotResponse{}, err
	}

	return oapi.SpotResponse{
		Id:          createdSpot.ID,
		Name:        spot.Name,
		Description: spot.Description,
		Address:     spot.Address,
		CreatedAt:   createdSpot.CreatedAt.Time.UTC(),
	}, nil
}

func (r *postgresSpotRepository) GetSpotByID(ctx context.Context, spotId uuid.UUID) (oapi.SpotResponse, error) {
	row, err := r.q.GetSpot(ctx, spotId)
	if err != nil {
		return oapi.SpotResponse{}, err
	}
	return oapi.SpotResponse{
		Id:          row.ID,
		Name:        row.Name,
		Description: row.Description,
		Address:     row.Address,
		CreatedAt:   row.CreatedAt.Time.UTC(),
	}, nil
}

func (r *postgresSpotRepository) UpdateSpotByID(ctx context.Context, spotId uuid.UUID, spot *oapi.SpotUpdate) (oapi.SpotResponse, error) {
	newSpot, err := r.q.GetSpot(ctx, spotId)
	if err != nil {
		return oapi.SpotResponse{}, err
	}

	if newName, err := spot.Name.Get(); err == nil {
		newSpot.Name = newName
	}
	if newDescription, err := spot.Description.Get(); err == nil {
		newSpot.Description = newDescription
	}
	if newAddress, err := spot.Address.Get(); err == nil {
		newSpot.Address = newAddress
	}
	params := sqlc.UpdateSpotParams{
		ID:              spotId,
		Name:            newSpot.Name,
		Description:     newSpot.Description,
		Address:         newSpot.Address,
		EmbeddingOpenai: newSpot.EmbeddingOpenai, // 変更なし（サボり？）
	}
	updated, err := r.q.UpdateSpot(ctx, params)
	if err != nil {
		return oapi.SpotResponse{}, err
	}
	return oapi.SpotResponse{
		Id:          updated.ID,
		Name:        updated.Name,
		Description: updated.Description,
		Address:     updated.Address,
		CreatedAt:   updated.CreatedAt.Time.UTC(),
	}, nil
}

func (r *postgresSpotRepository) SearchSpots(ctx context.Context, query string) ([]oapi.SpotResponse, error) {
	// LIKE句のためにワイルドカードを追加
	searchQuery := "%" + query + "%"
	rows, err := r.q.SearchSpots(ctx, searchQuery)
	if err != nil {
		return nil, err
	}

	spots := make([]oapi.SpotResponse, len(rows))
	for i, row := range rows {
		spots[i] = oapi.SpotResponse{
			Id:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			Address:     row.Address,
			CreatedAt:   row.CreatedAt.Time.UTC(),
		}
	}
	return spots, nil
}

func (r *postgresSpotRepository) SearchSpotsByEmbedding(ctx context.Context, embedding []float32) ([]oapi.SpotResponse, error) {
	vec := pgvector.NewVector(embedding)
	rows, err := r.q.SearchSpotsByEmbedding(ctx, &vec)
	if err != nil {
		return nil, err
	}

	spots := make([]oapi.SpotResponse, len(rows))
	for i, row := range rows {
		spots[i] = oapi.SpotResponse{
			Id:          row.ID,
			Name:        row.Name,
			Description: row.Description,
			Address:     row.Address,
			CreatedAt:   row.CreatedAt.Time.UTC(),
		}
	}
	return spots, nil
}
