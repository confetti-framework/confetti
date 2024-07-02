package ping

import (
    "github.com/matryer/is"
    "io"
    "net/http"
    "net/http/httptest"
    "src/app/http/controllers"
    "src/app/routes"
    "testing"
)

func Test_show_unkown_endpoint(t *testing.T) {
    // Given
    i := is.New(t)
    request := httptest.NewRequest(http.MethodGet, "/ping", nil)
    response := httptest.NewRecorder()

    // When
	routes.HandleApiRoute(response, request, controllers.Ping)
    result := response.Result()

    // Then
    i.Equal(result.StatusCode, http.StatusOK)
    i.Equal(getBody(result), "pong")
}

func getBody(result *http.Response) string {
    defer result.Body.Close()
    body, err := io.ReadAll(result.Body)
    if err != nil {
        panic(err)
    }
    return string(body)
}
