package controllers

import (
	"github.com/confetti-framework/framework/foundation/http/outcome"
	"github.com/confetti-framework/framework/inter"
	net "net/http"
)

// Ping is an endpoint with which you can check whether your application is responding.
func Ping(_ inter.Request) inter.Response {
	return outcome.Html("pong").Status(net.StatusOK)
}
