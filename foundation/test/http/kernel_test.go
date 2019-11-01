package test

import (
	"github.com/stretchr/testify/assert"
	"lanvard/src/app/http"
	"lanvard/src/bootstrap"
	net "net/http"
	"net/url"
	"testing"
)

func Test_handleRouting(t *testing.T) {
	app := bootstrap.App()

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
