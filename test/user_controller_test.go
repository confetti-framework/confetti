package test

import (
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/method"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_get_all_users(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/api/users",
	})

	ResponseByRequest(request)

	// @todo implement response
	// assert.Equal(t, outcome.Json("{\"test all users\"}"), result.Content())
}

func Test_get_user_by_id(t *testing.T) {
	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/api/user/64564",
	})

	result := ResponseByRequest(request)

	assert.Equal(t, "{\"email\":\"test@lanvard.com\",\"id\":64564}", result.Content())
}
