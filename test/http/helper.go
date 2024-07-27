package http

import (
	"context"
	"errors"
	"io"
	"net/http"
	"src/app/entity"
	"src/app/http/middleware"
)

type AuthServiceMock struct {
	Permissions []entity.Permission
}

func (a AuthServiceMock) Can(checkPermissions ...string) error {
	for _, permission := range checkPermissions {
		if !a.hasPermission(permission) {
			return errors.New("your permission does not have the required privileges. Permission: " + permission)
		}
	}
	return nil
}

func (a AuthServiceMock) hasPermission(checkPermission string) bool {
	for _, permission := range a.Permissions {
		if checkPermission == permission.Id {
			return true
		}
	}
	return false
}

func Auth(request *http.Request, permissions []entity.Permission) *http.Request {
	ctx := request.Context()
	ctx = context.WithValue(ctx, middleware.AuthServiceKey, AuthServiceMock{Permissions: permissions})
	return request.WithContext(ctx)
}

func GetRoute(routes []entity.Route, pattern string) entity.Route {
	for _, route := range routes {
		if route.Pattern == pattern {
			return route
		}
	}
	panic(`Route with pattern ` + pattern + ` not found`)
}

func GetBody(result *http.Response) string {
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
