package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/s1f10230101/INIAD_Team_Project_Group9Team3/oapi"
)

// 旅行プランの生成
// (POST /plans)
func (s *server) GeneratePlan(ctx context.Context, request oapi.GeneratePlanRequestObject) (oapi.GeneratePlanResponseObject, error) {
	// 1. PythonサービスのURL
	url := "http://python-backend:8000/generate-plan"

	// 2. Pythonサービスに送信するリクエストボディを作成
	pyReqBody := map[string]string{"prompt": request.Body.Prompt}
	jsonBody, err := json.Marshal(pyReqBody)
	if err != nil {
		return oapi.GeneratePlan500JSONResponse{Message: "Failed to create request body for python service"}, err
	}

	// 3. PythonサービスにHTTP POSTリクエストを送信
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return oapi.GeneratePlan500JSONResponse{Message: fmt.Sprintf("Failed to call python service: %v", err)}, err
	}

	if resp.StatusCode != http.StatusOK {
		// エラーの場合はボディを読んでメッセージを返す
		defer resp.Body.Close()
		return oapi.GeneratePlan500JSONResponse{Message: fmt.Sprintf("Python service returned error code: %d", resp.StatusCode)}, nil
	}

	// 4. Pythonサービスからのストリーミングレスポンスをそのままクライアントに返す
	// oapi.GeneratePlan200TexteventStreamResponse は io.Reader を Body として受け取るため、
	// http.Response.Body を直接渡すことで、ストリームを中継できる。
	// レスポンスボディのクローズは、レスポンスを処理するnet/httpサーバーに委ねられる。
	return oapi.GeneratePlan200TexteventStreamResponse{
		Body:          resp.Body,
		ContentLength: resp.ContentLength, // ContentLengthが設定されていればそれも渡す
	}, nil
}
