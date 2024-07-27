package ping

import (
	"github.com/matryer/is"
	net "net/http"
	"net/http/httptest"
	"src/app/entity"
	"src/app/http/route"
	helper "src/test/http"
	"testing"
)

var statusIndexMe = helper.GetRoute(route.Api, "GET /status")

func Test_index_status_but_unauthorized(t *testing.T) {
	// Given
	i := is.New(t)
	request := httptest.NewRequest(net.MethodGet, "/status", nil)
	request = helper.Auth(request, []entity.Permission{{Id: "/org1/project1/status/index"}})
	response := httptest.NewRecorder()

	// When
	route.HandleApiRoute(response, request, statusIndexMe)
	result := response.Result()

	// Then
	i.Equal(result.StatusCode, net.StatusUnauthorized)
}

func Test_index_status_authorized(t *testing.T) {
	// Given
	i := is.New(t)
	request := httptest.NewRequest(net.MethodGet, "/status", nil)
	request = helper.Auth(request, []entity.Permission{{Id: "/status/index"}})
	response := httptest.NewRecorder()

	// When
	route.HandleApiRoute(response, request, statusIndexMe)
	result := response.Result()

	// Then
	i.Equal(result.StatusCode, net.StatusOK)
}
