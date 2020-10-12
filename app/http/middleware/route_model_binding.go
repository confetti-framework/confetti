package middleware

import (
	"github.com/lanvard/contract/inter"
)

type RouteModelBinding struct{}

func (b RouteModelBinding) Handle(request inter.Request, next inter.Next) inter.Response {
	// request.App().Bind("user", func() model.User {
	// 	return model.User.Find(request.Parameter("user"))
	// })

	return next(request)
}
