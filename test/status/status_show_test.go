package ping

import (
    "github.com/matryer/is"
    "net/http"
    "net/http/httptest"
    "src/app/entity"
    "src/app/http/route"
    "src/test"
    "testing"
)

var statusIndexMe = test.GetRoute(route.Api, "/status")

func Test_index_status_but_unauthorized(t *testing.T) {
    // Given
    i := is.New(t)
    request := httptest.NewRequest(http.MethodGet, "/status", nil)
    request = test.Auth(request, []entity.Permission{{Id: "/org1/project1/status/index"}})
    response := httptest.NewRecorder()

    // When
    route.HandleApiRoute(response, request, statusIndexMe)
    result := response.Result()

    // Then
    i.Equal(result.StatusCode, http.StatusUnauthorized)
}

func Test_index_status_authorized(t *testing.T) {
    // Given
    i := is.New(t)
    request := httptest.NewRequest(http.MethodGet, "/status", nil)
    request = test.Auth(request, []entity.Permission{{Id: "/org1/project1/status/index"}})
    response := httptest.NewRecorder()

    // When
    route.HandleApiRoute(response, request, statusIndexMe)
    result := response.Result()

    // Then
    i.Equal(result.StatusCode, http.StatusOK)
}
