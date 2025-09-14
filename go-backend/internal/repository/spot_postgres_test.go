//go:build integration

package repository_test

import (
	"context"
	"log"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

var testPool *pgxpool.Pool

// TestMain はテストのセットアップとティアダウンを行います。
func TestMain(m *testing.M) {
	// compose.yamlで定義したDBに接続
	dbURL := "postgres://app_user:password@localhost:5432/app_db?sslmode=disable"
	var err error
	testPool, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("database接続失敗: %v", err)
	}
	defer testPool.Close()

	mg, err := migrate.New(
		"file://../../db/migration",
		dbURL,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}
	mg.Up()

	// Run tests
	m.Run()
}

func TestPostgresSpotRepository_CreateAndGetSpot(t *testing.T) {
	// --- Setup: トランザクションを開始 ---
	tx, err := testPool.Begin(context.Background())
	if err != nil {
		t.Fatalf("トランザクション開始失敗: %v", err)
	}
	// テスト終了時にロールバックすることで、DBの状態を元に戻す
	defer tx.Rollback(context.Background())

	// トランザクションを使ってリポジトリを初期化
	repo := repository.NewPostgresPostRepository(tx)

	// --- Test: CreateSpot ---
	input := &oapi.SpotResister{
		Name:        "テスト用観光地 (Postgres)",
		Description: "これはインテグレーションテスト用のデータです。",
		Address:     "東京都テスト区",
	}

	createdSpot, err := repo.CreateSpot(context.Background(), input)
	if err != nil {
		t.Fatalf("CreateSpot失敗: %v", err)
	}

	// --- Assert: CreateSpotの結果を検証 ---
	if createdSpot.Id == uuid.Nil {
		t.Fatalf("CreateSpotでIDが設定されていない%d\n", createdSpot.Id)
	}
	if createdSpot.Name != input.Name {
		t.Fatalf("CreateSpotでNameが正しく設定されていない: got %s, want %s", createdSpot.Name, input.Name)
	}
	if createdSpot.Description != input.Description {
		t.Fatalf("CreateSpotでDescriptionが正しく設定されていない: got %s, want %s", createdSpot.Description, input.Description)
	}
	if createdSpot.Address != input.Address {
		t.Fatalf("CreateSpotでAddressが正しく設定されていない: got %s, want %s", createdSpot.Address, input.Address)
	}
	if createdSpot.CreatedAt.IsZero() {
		t.Fatalf("CreateSpotでCreatedAtが設定されていない: got %v", createdSpot.CreatedAt)
	}
	// --- Test: GetSpotByID (同じトランザクション内で取得) ---
	fetchedSpot, err := repo.GetSpotByID(context.Background(), createdSpot.Id)
	if err != nil {
		t.Fatalf("GetSpotByID失敗: %v", err)
	}

	// --- Assert: GetSpotByIDの結果を検証 ---
	if createdSpot.Id != fetchedSpot.Id {
		t.Fatalf("GetSpotByIDでIDが一致しない: got %d, want %d", fetchedSpot.Id, createdSpot.Id)
	}
	if createdSpot.Name != fetchedSpot.Name {
		t.Fatalf("GetSpotByIDでNameが一致しない: got %s, want %s", fetchedSpot.Name, createdSpot.Name)
	}
	if createdSpot.Description != fetchedSpot.Description {
		t.Fatalf("GetSpotByIDでDescriptionが一致しない: got %s, want %s", fetchedSpot.Description, createdSpot.Description)
	}
}
