package controllers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/outcome"
	net "net/http"
)

func Ping(_ inter.Request) inter.Response {
	return outcome.Html("pong").Status(net.StatusOK)
}
