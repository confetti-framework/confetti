package middleware

import (
	"net/http"
	"src/internal/pkg/handler"
)

type AuthMiddleware struct {
	Permission string
}

func (a AuthMiddleware) Handle(next handler.Controller) handler.Controller {
	return func(w http.ResponseWriter, req *http.Request) error {

		// Here you can do something before the request
		// For example, check if the user has the permission

		err := next(w, req)

		// Here you can do something after the request
		// For example, log the response status

		return err
	}
}
