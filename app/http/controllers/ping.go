package controllers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/outcome"
	net "net/http"
)

// Ping is an endpoint with which you can check whether your application is responding.
func Ping(_ inter.Request) inter.Response {
	return outcome.Html("pong").Status(net.StatusOK)
}
