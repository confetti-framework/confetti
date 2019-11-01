package test

import (
	"github.com/stretchr/testify/assert"
	"laravelgo/foundation"
	"laravelgo/src/app/http"
	net "net/http"
	"net/url"
	"testing"
)

func Test_handleRouting(t *testing.T) {
	app := foundation.Application()

	kernel := http.Kernel(app)
	
	request := net.Request{
		Method:     "GET",
		Host:       "example.com",
		URL:        &url.URL{Path: "/test?123"},
		Header:     net.Header{},
		RequestURI: "/",
	}
	response := kernel.Handle(request)

	assert.Equal(t, "response", response)
}