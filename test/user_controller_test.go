package test

import (
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/method"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_get_all_users(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/api/users",
	})

	response := ResponseByRequest(request)

	assert.Equal(t, outcome.Json("{\"test all users\"}"), response.Body())
}

func Test_get_user_by_id(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/api/user/64564",
	})

	result := ResponseByRequest(request)

	assert.Equal(t, "{\"email\":\"test@lanvard.com\",\"id\":64564}", result.Body())
}
