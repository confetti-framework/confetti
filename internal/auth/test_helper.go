package auth

import (
	"context"
	"errors"
	"net/http"
)

type ServiceMock struct {
	Permissions []Permission
}

func (a ServiceMock) Can(checkPermissions ...string) error {
	for _, permission := range checkPermissions {
		if !a.hasPermission(permission) {
			return errors.New("You do not have the required privileges for permission: " + permission)
		}
	}
	return nil
}

func (a ServiceMock) hasPermission(checkPermission string) bool {
	for _, permission := range a.Permissions {
		if checkPermission == permission.Id {
			return true
		}
	}
	return false
}

func Auth(request *http.Request, permissions []Permission) *http.Request {
	ctx := request.Context()
	ctx = context.WithValue(ctx, ServiceKey, ServiceMock{Permissions: permissions})
	return request.WithContext(ctx)
}
