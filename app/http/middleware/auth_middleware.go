package middleware

import (
    "context"
    "errors"
    "net/http"
    "path"
    "src/app/config"
    "src/app/entity"
    service "src/app/service"
    "strings"
)

const AuthServiceKey = "auth_service"

type AuthServiceInterface interface {
    Can(checkPermissions ...string) error
}

type auth struct {
    permissions []string
}

func Auth(permissions ...string) entity.Middleware {
    result := []string{}
    for _, permission := range permissions {
        if !strings.HasPrefix(permission, "/") {
            // Prefix current project. E.g. `status/index` to `/org1/rep1/status/index`
            permission = "/" + path.Join(config.AppInfo.Repository, config.AppInfo.Service, permission)
        }
        result = append(result, permission)
    }
    return auth{permissions: result}
}

// Handle authenticates users
func (a auth) Handle(next entity.Controller) entity.Controller {
    return func(w http.ResponseWriter, r *http.Request) error {
        ctx := r.Context()
        authService := ctx.Value(AuthServiceKey)
        if authService == nil {
            authService = service.AuthService{}.InitByRequest(r)
            ctx = context.WithValue(ctx, AuthServiceKey, authService)
            r.WithContext(ctx)
        }

        if err := authService.(AuthServiceInterface).Can(a.permissions...); err == nil {
            return next(w, r)
        } else {
            return errors.Join(err, entity.UserError{HttpStatus: http.StatusUnauthorized})
        }
    }
}
