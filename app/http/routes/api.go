package routes

import (
    "encoding/json"
    "errors"
    "net/http"
    "src/app/http/controllers"
    "src/app/http/entities"
    "src/app/http/middlewares"
)

type HttpStatusGetter interface {
    GetHttpStatus() int
}

var Api = []entities.Route{
    route("/ping", controllers.Ping),
    route("/status", controllers.Status).AppendMiddleware(middlewares.Auth),
}

func HandleApiRoute(response http.ResponseWriter, request *http.Request, route entities.Route) {
    handler := entities.MiddlewareChain(route.GetMiddlewares())

    err := handler(route.Controller)(response, request)
    if err != nil {
        apiErrorHandler(response, err)
    }
}

func route(pattern string, controller entities.Controller) entities.Route {
    return entities.Route{Pattern: pattern, Controller: controller}
}

// apiErrorHandler converts the error to a reponse. Feel free to modify this to your needs.
func apiErrorHandler(writer http.ResponseWriter, err error) {
    status := getStatusCodeFromError(err)

    // Set the response header to application/json
    writer.Header().Set("Content-Type", "application/json")
    // Write the status code as 500 Internal Server Error
    writer.WriteHeader(status)
    // Create a JSON response with the error message
    _ = json.NewEncoder(writer).Encode(map[string]string{"message": err.Error()})
}

func getStatusCodeFromError(err error) int {
    if err, ok := err.(HttpStatusGetter); ok {
        return err.GetHttpStatus()
    }

    // Check for unwrapped errors
    if unwrappedErr := errors.Unwrap(err); unwrappedErr != nil {
        return getStatusCodeFromError(unwrappedErr)
    }

    return http.StatusInternalServerError
}
