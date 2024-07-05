package ping

import (
    "github.com/matryer/is"
    "net/http"
    "net/http/httptest"
    "src/app/http/routes"
    "src/test"
    "testing"
)

var statusIndexMe = test.GetRoute(routes.Api, "/status")

func Test_index_status_but_unauthorized(t *testing.T) {
    // Given
    i := is.New(t)
    request := httptest.NewRequest(http.MethodGet, "/status", nil)
    response := httptest.NewRecorder()

    // When
    routes.HandleApiRoute(response, request, statusIndexMe)
    result := response.Result()

    // Then
    i.Equal(result.StatusCode, http.StatusUnauthorized)
}
