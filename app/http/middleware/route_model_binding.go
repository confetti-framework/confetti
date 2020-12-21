package middleware

import (
	"github.com/confetti-framework/contract/inter"
)

type RouteModelBinding struct{}

func (b RouteModelBinding) Handle(request inter.Request, next inter.Next) inter.Response {

	// Feel free to bind your models here.
	// request.App().Bind("user", func() model.User {
	// 	return model.User.Find(request.Parameter("user"))
	// })

	return next(request)
}
