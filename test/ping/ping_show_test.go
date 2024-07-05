package ping

import (
    "github.com/matryer/is"
    "net/http"
    "net/http/httptest"
    routes "src/app/http/routes"
    "src/test"
    "testing"
)

var pingShow = test.GetRoute(routes.Api, "/ping")

func Test_show_ping_endpoint(t *testing.T) {
    // Given
    i := is.New(t)
    request := httptest.NewRequest(http.MethodGet, "/ping", nil)
    response := httptest.NewRecorder()

    // When
    routes.HandleApiRoute(response, request, pingShow)
    result := response.Result()

    // Then
    i.Equal(result.StatusCode, http.StatusOK)
    i.Equal(test.GetBody(result), "pong")
}
