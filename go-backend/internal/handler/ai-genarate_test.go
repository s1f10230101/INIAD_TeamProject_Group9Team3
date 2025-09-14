package handler_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

func TestGeneratePlan(t *testing.T) {
	// --- setup ---
	router := setupTestRouter()

	// --- execute ---
	// 1. 旅行プラン生成のリクエストを作成
	prompt := "家族で楽しめる沖縄旅行"
	body, err := json.Marshal(oapi.Prompt{Prompt: prompt})
	if err != nil {
		t.Fatalf("リクエストボディの作成に失敗しました: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/v1/plans", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	// ルーターにリクエストを送信
	router.ServeHTTP(rec, req)

	// --- verify ---
	// レスポンスコードが200 OKであることを確認
	if rec.Code != http.StatusOK {
		t.Fatalf("期待されるステータスコード %d ではありません: got %d", http.StatusOK, rec.Code)
	}

	// レスポンスヘッダーが text/event-stream であることを確認
	expectedContentType := "text/event-stream"
	contentType := rec.Header().Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("期待されるContent-Type %q ではありません: got %q", expectedContentType, contentType)
	}

	// ストリームの内容を読み取る
	reader := bufio.NewReader(rec.Body)
	var receivedMessages []string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("ストリームの読み取りに失敗しました: %v", err)
		}
		// SSEの "data: " プレフィックスを削除し、前後の空白をトリム
		trimmedLine := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		// 空行は無視
		if trimmedLine == "" {
			continue
		}
		receivedMessages = append(receivedMessages, trimmedLine)
		println("Received message:", trimmedLine)
	}

	// AIGenerateFake が生成するメッセージの期待値
	expectedMessages := []string{
		"{\"text\": \"これはテストメッセージ カウント 1\"}",
		"{\"text\": \"これはテストメッセージ カウント 2\"}",
		"{\"text\": \"これはテストメッセージ カウント 3\"}",
		"{\"text\": \"prompt: 家族で楽しめる沖縄旅行\"}",
		"event: done",
	}

	if len(receivedMessages) != len(expectedMessages) {
		t.Fatalf("期待されるメッセージ数 %d ではありません: got %d. Received: %v", len(expectedMessages), len(receivedMessages), receivedMessages)
	}

	for i, expected := range expectedMessages {
		if receivedMessages[i] != expected {
			t.Errorf("期待されるメッセージ %q at index %d ではありません: got %q", expected, i, receivedMessages[i])
		}
	}
}
