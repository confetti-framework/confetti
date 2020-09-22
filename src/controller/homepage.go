package controller

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing/outcome"
)

func Homepage(request inter.Request) inter.Response {
	return outcome.Html("ok")
}
