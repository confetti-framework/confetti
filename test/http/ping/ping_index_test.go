package ping

import (
	"net/http"
	"net/http/httptest"
	"src/cmd/api/command"
	"src/internal/pkg/handler"
	"testing"
)

var pingIndex = handler.GetRoute(command.ApiRoutes, "GET /ping")

func Test_show_ping_endpoint(t *testing.T) {
	// Given
	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	response := httptest.NewRecorder()

	// When
	handler.HandleApiRoute(response, request, pingIndex)
	result := response.Result()

	// Then
	if result.StatusCode != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, result.StatusCode)
	}
	if body := handler.GetBody(result); body != "pong" {
		t.Errorf("expected body %v, got %v", "pong", body)
	}
}
