package repository_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/stretchr/testify/assert"
)

var testPool *pgxpool.Pool

func TestMain(m *testing.M) {
	// Setup
	// compose.yamlで定義したDBに接続
	databaseUrl := "postgres://app_user:password@localhost:5432/app_db?sslmode=disable"
	var err error
	testPool, err = pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer testPool.Close()

	// Run tests
	code := m.Run()

	// Teardown
	os.Exit(code)
}

func TestPostgresSpotRepository_CreateAndGetSpot(t *testing.T) {
	// --- Setup: トランザクションを開始 ---
	tx, err := testPool.Begin(context.Background())
	assert.NoError(t, err)
	// テスト終了時にロールバックすることで、DBの状態を元に戻す
	defer tx.Rollback(context.Background())

	// トランザクションを使ってリポジトリを初期化
	repo := repository.NewPostgresPostRepositoryForTest(tx)

	// --- Test: CreateSpot ---
	input := &api.SpotInput{
		Name:        "テスト用観光地 (Postgres)",
		Description: "これはインテグレーションテスト用のデータです。",
		Address:     "東京都テスト区",
	}

	createdSpot, err := repo.CreateSpot(input)
	assert.NoError(t, err)

	// --- Assert: CreateSpotの結果を検証 ---
	assert.NotEmpty(t, createdSpot.Id)
	assert.Equal(t, input.Name, createdSpot.Name)
	assert.Equal(t, input.Description, *createdSpot.Description)
	assert.Equal(t, input.Address, *createdSpot.Address)
	assert.NotZero(t, createdSpot.CreatedAt)

	// --- Test: GetSpotByID (同じトランザクション内で取得) ---
	fetchedSpot, err := repo.GetSpotByID(createdSpot.Id)
	assert.NoError(t, err)

	// --- Assert: GetSpotByIDの結果を検証 ---
	assert.Equal(t, createdSpot.Id, fetchedSpot.Id)
	assert.Equal(t, createdSpot.Name, fetchedSpot.Name)
	assert.Equal(t, *createdSpot.Description, *fetchedSpot.Description)
}
