package ping

import (
    "github.com/matryer/is"
    "net/http"
    "net/http/httptest"
    "src/app/http/routes"
    "src/test"
    "testing"
)

var userShowMe = test.GetRoute(routes.Api, "/users/me")

func Test_show_user_me_but_unauthorized(t *testing.T) {
    // Given
    i := is.New(t)
    request := httptest.NewRequest(http.MethodGet, "/users/me", nil)
    response := httptest.NewRecorder()

    // When
    routes.HandleApiRoute(response, request, userShowMe)
    result := response.Result()

    // Then
    i.Equal(result.StatusCode, http.StatusUnauthorized)
}
