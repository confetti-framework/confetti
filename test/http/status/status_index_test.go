package ping

import (
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"src/cmd/api/command"
	"src/internal/auth"
	"src/internal/pkg/handler"
	"testing"
)

var statusIndexMe = handler.GetRoute(command.ApiRoutes, "GET /status")

func Test_index_status_but_unauthorized(t *testing.T) {
	// Given
	i := is.New(t)
	request := httptest.NewRequest(http.MethodGet, "/status", nil)
	request = auth.MockRequest(request, []auth.Permission{{Id: "/org1/project1/status/index"}})
	response := httptest.NewRecorder()

	// When
	handler.HandleApiRoute(response, request, statusIndexMe)
	result := response.Result()

	// Then
	i.Equal(result.StatusCode, http.StatusUnauthorized)
}

func Test_index_status_authorized(t *testing.T) {
	// Given
	i := is.New(t)
	request := httptest.NewRequest(http.MethodGet, "/status", nil)
	request = auth.MockRequest(request, []auth.Permission{{Id: "/status/index"}})
	response := httptest.NewRecorder()

	// When
	handler.HandleApiRoute(response, request, statusIndexMe)
	result := response.Result()

	// Then
	i.Equal(result.StatusCode, http.StatusOK)
}

func Test_index_status_cpu(t *testing.T) {
	// Given
	i := is.New(t)
	request := httptest.NewRequest(http.MethodGet, "/status", nil)
	request = auth.MockRequest(request, []auth.Permission{{Id: "/status/index"}})
	response := httptest.NewRecorder()

	// When
	handler.HandleApiRoute(response, request, statusIndexMe)
	result := response.Result()
	body := handler.GetBody(response.Result())

	// Then
	i.Equal(result.StatusCode, http.StatusOK)
	i.Equal(handler.GetField(body, "status"), "active")
}
