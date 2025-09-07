package repository

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/api"
	//"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/internal/repository"
)

func TestCreateSpot(t *testing.T) {
	repo := NewSpotRepositoryInmemory()
	inputSpot := &api.SpotInput{
		Name:        "テスト用観光",
		Description: "これはテスト用の説明です。",
		Address:     "東京都テスト区1-2-3",
	}

	_, err := repo.CreateSpot(context.Background(), inputSpot)

	if err != nil {
		t.Fatalf("CreateSpot failed, expected no error, but got: %v", err)
	}

	if len(repo.postsDB) != 1 {
		t.Fatalf("Expected postsDB to have 1 spot, but got: %d", len(repo.postsDB))
	}

	var savedSpot api.Spot
	for _, spot := range repo.postsDB {
		savedSpot = spot
		break
	}

	if savedSpot.Name != inputSpot.Name {
		t.Errorf("Expected spot name to be '%s', but got '%s'", inputSpot.Name, savedSpot.Name)
	}
	if savedSpot.Description != &inputSpot.Description {
		t.Errorf("Expected spot description to be '%s', but got '%v'", inputSpot.Description, savedSpot.Description)
	}
	if savedSpot.Address != &inputSpot.Address {
		t.Errorf("Expected spot address to be '%s', but got '%v'", inputSpot.Address, savedSpot.Address)
	}

	// IDが自動生成されているか（空でないか）を確認
	if savedSpot.Id.String() == "00000000-0000-0000-0000-000000000000" {
		t.Error("Expected spot ID to be generated, but it was a zero UUID")
	}

	// CreatedAtが設定されているか（ゼロ値でないか）を確認
	if savedSpot.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set, but it was a zero time")
	}
}

func TestGetAllSpots(t *testing.T) {
	t.Run("Success: Empty slice for empty repository", func(t *testing.T) {
		// 準備
		repo := NewSpotRepositoryInmemory()

		// 実行
		spots, err := repo.GetAllSpots(context.Background())

		// 検証
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}
		if spots == nil {
			t.Fatal("expected an empty slice, but got nil")
		}
		if len(spots) != 0 {
			t.Errorf("expected 0 spots, but got %d", len(spots))
		}
	})

	t.Run("Success: Returns all spots", func(t *testing.T) {
		// 準備
		repo := NewSpotRepositoryInmemory()
		spot1 := api.Spot{Id: uuid.New(), Name: "観光地1", CreatedAt: time.Now()}
		spot2 := api.Spot{Id: uuid.New(), Name: "観光地2", CreatedAt: time.Now()}
		repo.postsDB[spot1.Id] = spot1
		repo.postsDB[spot2.Id] = spot2

		// 実行
		spots, err := repo.GetAllSpots(context.Background())

		// 検証
		if err != nil {
			t.Fatalf("expected no error, but got %v", err)
		}
		if len(spots) != 2 {
			t.Errorf("expected 2 spots, but got %d", len(spots))
		}
	})
}

func TestNewPostRepositoryInmemory(t *testing.T) {
	repo := NewSpotRepositoryInmemory()

	if repo == nil {
		t.Fatal("NewPostRepositoryInmemory() returned nil")
	}
	if repo.postsDB == nil {
		t.Fatal("postsDB map was not initialized")
	}
}

func TestGetSpotByID(t *testing.T) {
	repo := NewSpotRepositoryInmemory()
	preloadedSpot := api.Spot{
		Id:        uuid.New(),
		Name:      "既存の観光地",
		CreatedAt: time.Now(),
	}

	repo.postsDB[preloadedSpot.Id] = preloadedSpot

	t.Run("Success: Spot found", func(t *testing.T) {
		gotSpot, err := repo.GetSpotByID(context.Background(), preloadedSpot.Id)
		if err != nil {
			t.Fatalf("expected no error, but got: %v", err)
		}
		if gotSpot.Id != preloadedSpot.Id {
			t.Errorf("expected spot ID %v, but got %v", preloadedSpot.Id, gotSpot.Id)
		}
	})

	t.Run("Failure: Spot not found", func(t *testing.T) {
		// 実行 & 検証
		nonExistentID := uuid.New()
		_, err := repo.GetSpotByID(context.Background(), nonExistentID)
		if err == nil {
			t.Fatal("expected an error for non-existent ID, but got nil")
		}
	})
}

func TestUpdateSpotByID(t *testing.T) {
	repo := NewSpotRepositoryInmemory()

	description := "更新前の説明"
	address := "更新前の住所"

	initialSpot := api.Spot{
		Id:          uuid.New(),
		Name:        "更新前の名前",
		Description: &description,
		Address:     &address,
		CreatedAt:   time.Now(),
	}

	repo.postsDB[initialSpot.Id] = initialSpot

	updateInput := &api.SpotInput{
		Name:        "更新後の名前",
		Description: "更新後の説明",
		Address:     "更新後の住所",
	}

	t.Run("Success: Spot updated", func(t *testing.T) {
		// 実行
		updatedSpot, err := repo.UpdateSpotByID(context.Background(), initialSpot.Id, updateInput)
		if err != nil {
			t.Fatalf("expected no error, but got: %v", err)
		}

		// 検証
		if updatedSpot.Name != updateInput.Name {
			t.Errorf("expected name to be '%s', but got '%s'", updateInput.Name, updatedSpot.Name)
		}
		if updatedSpot.Description != &updateInput.Description {
			t.Errorf("expected description to be '%s', but got '%v'", updateInput.Description, updatedSpot.Description)
		}
		// IDとCreatedAtが変更されていないことを確認
		if updatedSpot.Id != initialSpot.Id {
			t.Error("ID should not be changed on update")
		}
		if !updatedSpot.CreatedAt.Equal(initialSpot.CreatedAt) {
			t.Error("CreatedAt should not be changed on update")
		}
	})

	t.Run("Failure: Spot not found", func(t *testing.T) {
		// 実行 & 検証
		nonExistentID := uuid.New()
		_, err := repo.UpdateSpotByID(context.Background(), nonExistentID, updateInput)
		if err == nil {
			t.Fatal("expected an error for non-existent ID, but got nil")
		}
	})
}
