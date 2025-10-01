package usecase

import (
	"fmt"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// Exaple_buildPrompt は buildPrompt 関数のテストを行います。(外部サービス依存なし)
func Example_buildPrompt() {
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

	systemPrompt, userPromptResult, err := _buildPrompt(spots, userPrompt)
	if err != nil {
		panic(err)
	}

	// 結果を表示
	fmt.Println("System Prompt:")
	fmt.Println(systemPrompt)
	fmt.Println("User Prompt:")
	fmt.Println(userPromptResult)
	// Output:
	// System Prompt:
	//
	// あなたは旅行プランのプロです。以下の参考情報とユーザーの要望を元に、魅力的な旅行プランを提案してください。
	// 参考情報:
	// - 名前: 東京タワー
	//   説明: 東京のシンボル的存在の展望塔。
	//   住所: 東京都港区芝公園4丁目2-8
	// - 名前: 浅草寺
	//   説明: 東京最古の寺院で、雷門が有名。
	//   住所: 東京都台東区浅草2丁目3-1
	//
	// User Prompt:
	//
	// 家族で楽しめる東京旅行
}
