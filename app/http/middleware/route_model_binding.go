package middleware

import (
	"github.com/confetti-framework/contract/inter"
)

type RouteModelBinding struct{}

func (b RouteModelBinding) Handle(request inter.Request, next inter.Next) inter.Response {

	// Feel free to bind your models here.
	// request.App().Singleton("user", func() model.User {
	// 	userId := request.Parameter("user_id").Int()
	// 	return model.FindUser(userId)
	// })

	return next(request)
}
