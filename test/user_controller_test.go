package test

import (
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/method"
	"github.com/lanvard/foundation/http/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_create_user(t *testing.T) {

	request := http.NewRequest(http.Options{
		Method: method.Get,
		Url:    "/users",
	})

	result := ResponseByRequest(request)

	assert.Equal(t, outcome.Json("{\"test all users\"}"), result)
}
