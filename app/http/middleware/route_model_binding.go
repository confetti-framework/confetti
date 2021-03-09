package middleware

import (
	"github.com/confetti-framework/contract/inter"
)

// RouteModelBinding For simple transformation of a parameter to a model, you can bind a
// model to a parameter.
type RouteModelBinding struct{}

// Handle For simple transformation of a parameter to a model, you can bind a
// model to a parameter.
func (b RouteModelBinding) Handle(request inter.Request, next inter.Next) inter.Response {

	// Feel free to bind your models here.
	// request.App().Singleton("user", func() model.User {
	// 	userId := request.Parameter("user_id").Int()
	// 	return model.FindUser(userId)
	// })

	return next(request)
}
