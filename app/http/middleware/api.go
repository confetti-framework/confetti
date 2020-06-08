package middleware

import (
	"github.com/lanvard/contract/inter"
)

type Api struct{}

func (a Api) Handle(request inter.Request, next inter.Next) inter.Response {
	return next(request)
}
