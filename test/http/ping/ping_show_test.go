package ping

import (
    "github.com/matryer/is"
    net "net/http"
    "net/http/httptest"
    routes "src/app/http/route"
    helper "src/test/http"
    "testing"
)

var pingShow = helper.GetRoute(routes.Api, "GET /ping")

func Test_show_ping_endpoint(t *testing.T) {
    // Given
    i := is.New(t)
    request := httptest.NewRequest(net.MethodGet, "/ping", nil)
    response := httptest.NewRecorder()

    // When
    routes.HandleApiRoute(response, request, pingShow)
    result := response.Result()

    // Then
    i.Equal(result.StatusCode, net.StatusOK)
    i.Equal(helper.GetBody(result), "pong")
}
