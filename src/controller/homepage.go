package controller

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/routing/outcome"
	net "net/http"
)

func Homepage(request inter.Request) inter.Response {
	return outcome.Html(errors.New("User not found").Status(net.StatusNotFound))
}
