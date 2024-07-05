package middlewares

import (
    "net/http"
    "src/app/http/entities"
)

type unauthorizedError struct {
    httpStatus int
}

func (e unauthorizedError) Error() string {
    return "user is unauthorized"
}

func (e unauthorizedError) GetHttpStatus() int {
    return e.httpStatus
}

// Auth authenticates users
func Auth(next entities.Controller) entities.Controller {
    return func(w http.ResponseWriter, r *http.Request) error {
        // Perform authentication logic here
        isAuthenticated := false // For demonstration purposes

        if isAuthenticated {
            return next(w, r)
        } else {
            return unauthorizedError{httpStatus: http.StatusUnauthorized}
        }
    }
}
