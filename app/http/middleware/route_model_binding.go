package middleware

import (
	"github.com/lanvard/contract/inter"
)

type RouteModelBinding struct{}

func (b RouteModelBinding) Handle(request inter.Request, next inter.Next) inter.Response {
	// request.App().Instance("user", func() model.User {
	// 	return model.User.Find(request.Value("user"))
	// })

	return next(request)
}
