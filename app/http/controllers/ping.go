package controllers

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/routing/outcome"
	net "net/http"
)

func Ping(_ inter.Request) inter.Response {
	return outcome.Html(errors.New("pong").Status(net.StatusOK))
}
