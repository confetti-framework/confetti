package route

import (
    "encoding/json"
    "errors"
    "net/http"
    "src/app/entity"
    "src/app/http/controller"
    "src/app/http/middleware"
)

type HttpStatusGetter interface {
    GetHttpStatus() int
}

var Api = []entity.Route{
    route("/ping", controllers.Ping),
    route("/status", controllers.Status).AppendMiddleware(middleware.Auth("status/index")),
}

func HandleApiRoute(response http.ResponseWriter, request *http.Request, route entity.Route) {
    handler := entity.MiddlewareChain{Middlewares: route.GetMiddlewares()}

    err := handler.Handle(route.Controller)(response, request)
    if err != nil {
        apiErrorHandler(response, err)
    }
}

func route(pattern string, controller entity.Controller) entity.Route {
    return entity.Route{Pattern: pattern, Controller: controller}
}

// apiErrorHandler converts the error to a reponse. Feel free to modify this to your needs.
func apiErrorHandler(writer http.ResponseWriter, err error) {
    status := http.StatusInternalServerError
    var detailErr entity.UnauthorizedError
    if errors.As(err, &detailErr) {
        status = detailErr.GetHttpStatus()
    }

    // Set the response header to application/json
    writer.Header().Set("Content-Type", "application/json")
    // Write the status code as 500 Internal Server Error
    writer.WriteHeader(status)
    // Create a JSON response with the error message
    _ = json.NewEncoder(writer).Encode(map[string]string{"message": err.Error()})
}
