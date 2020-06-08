package middleware

import (
	"github.com/lanvard/contract/inter"
)

type Web struct{}

func (a Web) Handle(request inter.Request, next inter.Next) inter.Response {
	return next(request)
}
