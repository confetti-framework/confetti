package status

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"src/cmd/api/command"
	"src/internal/pkg/handler"
	"testing"
)

var statusIndex = handler.GetRoute(command.ApiRoutes, "GET /status")

func Test_show_status_endpoint(t *testing.T) {
	// Given
	request := httptest.NewRequest(http.MethodGet, "/status", nil)
	response := httptest.NewRecorder()

	// When
	handler.HandleApiRoute(response, request, statusIndex)
	result := response.Result()

	// Then
	if result.StatusCode != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, result.StatusCode)
	}
	var responseBody map[string]string
	if err := json.NewDecoder(result.Body).Decode(&responseBody); err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}
	if status, ok := responseBody["status"]; !ok || status != "active" {
		t.Errorf("expected status %v, got %v", "active", status)
	}
}
