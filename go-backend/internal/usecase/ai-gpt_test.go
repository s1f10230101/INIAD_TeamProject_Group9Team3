package usecase

import (
	"strings"
	"testing"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

func TestBuildPrompt(t *testing.T) {
	// テスト用の観光地データ
	spots := []oapi.SpotResponse{
		{
			Name:        "東京タワー",
			Description: "東京のシンボル的存在の展望塔。",
			Address:     "東京都港区芝公園4丁目2-8",
		},
		{
			Name:        "浅草寺",
			Description: "東京最古の寺院で、雷門が有名。",
			Address:     "東京都台東区浅草2丁目3-1",
		},
	}

	userPrompt := "家族で楽しめる東京旅行"

	systemPrompt, userPromptResult, err := buildPrompt(spots, userPrompt)

	if err != nil {
		t.Fatalf("buildPromptでエラーが発生しました: %v", err)
	}
	// systemPromptに観光地情報が含まれていることを確認
	if !strings.Contains(systemPrompt, "東京タワー") || !strings.Contains(systemPrompt, "浅草寺") {
		t.Errorf("systemPromptに観光地情報が含まれていません: %s", systemPrompt)
	}
	t.Logf("Generated systemPrompt: %s", systemPrompt)

	// userPromptResultが正しいことを確認
	expectedUserPrompt := `
家族で楽しめる東京旅行
 `
	if userPromptResult != expectedUserPrompt {
		t.Errorf("userPromptResultが期待値と異なります: got %q, want %q", userPromptResult, expectedUserPrompt)
	}
	t.Logf("Generated userPrompt: %s", userPromptResult)
}
