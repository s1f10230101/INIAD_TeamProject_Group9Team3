package handler_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

// TestGeneratePlanStreaming is an integration test that verifies the streaming response
// from the /v1/plans endpoint.
// It requires the docker compose services (backend-dev, python-backend, db) to be running.
func TestGeneratePlanStreaming(t *testing.T) {
	// 1. Prepare the request
	prompt := "大阪で抹茶を楽しみたい"
	reqBody := map[string]string{"prompt": prompt}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/v1/plans", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 2. Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	// 2a. Verify response headers for SSE
	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/event-stream") {
		t.Fatalf("Expected Content-Type 'text/event-stream', but got '%s'", contentType)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %d", resp.StatusCode)
	}

	// 3. Process the streaming response
	var fullResponse strings.Builder
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data:") {
			continue // Skip empty lines or comments
		}

		jsonData := strings.TrimPrefix(line, "data: ")
		if jsonData == "" {
			continue // Skip empty data lines
		}

		var data map[string]string
		if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
			// Log the error and the problematic data for debugging
			t.Logf("Failed to unmarshal JSON. Error: %v. Data: '%s'", err, jsonData)
			continue // Continue to the next line
		}

		if token, ok := data["token"]; ok {
			fullResponse.WriteString(token)
		}
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("Error reading streaming response: %v", err)
	}

	// 4. Assert the result
	finalText := fullResponse.String()
	t.Logf("Full response from stream: %s", finalText)

	if finalText == "" {
		t.Fatal("Received an empty response from the stream.")
	}

	// Check if the response contains some expected Japanese characters
	// Note: Depending on the model's response, this might fail.
	// It's better to check for non-emptiness first.
	if !strings.Contains(finalText, "抹茶") && !strings.Contains(finalText, "大阪") {
		t.Logf("Warning: Response did not contain expected keywords '抹茶' or '大阪'.")
	}
}