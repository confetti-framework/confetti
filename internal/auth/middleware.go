package auth

import (
	"context"
	"net/http"
	"path"
	"src/config"
	"src/internal/pkg/handler"
	"strings"
)

const ServiceKey = "auth_service"

type ServiceInterface interface {
	Can(checkPermissions ...string) error
}

type middleware struct {
	permissions []string
}

func Middleware(permissions ...string) handler.Middleware {
	result := []string{}
	for _, permission := range permissions {
		if !strings.HasPrefix(permission, "/") {
			// Prefix current project. E.g. `status/index` to `/org1/rep1/status/index`
			permission = path.Join(config.AppInfo.Permission, permission)
		}
		result = append(result, permission)
	}
	return middleware{permissions: result}
}

// Handle authenticates users
func (a middleware) Handle(next handler.Controller) handler.Controller {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		authService := ctx.Value(ServiceKey)
		if authService == nil {
			authService = Service{}.InitByRequest(r)
			ctx = context.WithValue(ctx, ServiceKey, authService)
			r.WithContext(ctx)
		}

		if err := authService.(ServiceInterface).Can(a.permissions...); err != nil {
			return err
		}

		return next(w, r)
	}
}
