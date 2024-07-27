package route

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    net "net/http"
    "src/app/config"
    "src/app/entity"
    "src/app/http/controller"
    "src/app/http/middleware"
    "strings"
    "time"
)

type HttpStatusGetter interface {
    GetHttpStatus() int
}

var apiRoutes = []entity.Route{
    route("GET /ping", controller.Ping),
    route("GET /status", controller.Status).AppendMiddleware(middleware.Auth("status/index")),
}

func GetApiRoutes() []entity.Route {
    return appendApiByPath(apiRoutes)
}

func HandleApiRoute(response net.ResponseWriter, request *net.Request, route entity.Route) {
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
func apiErrorHandler(writer net.ResponseWriter, err error) {
    status := net.StatusInternalServerError
    publicMessage := ""
    internalMessage := ""
    report := fmt.Sprintf("%d", time.Now().UnixMilli())
    var systemError entity.SystemError
    var userError entity.UserError
    if errors.As(err, &systemError) {
        // Handle system error
        status = systemError.GetHttpStatus()
        publicMessage = "internal server error"
        internalMessage = err.Error()
    } else if errors.As(err, &userError) {
        // Handle user error
        status = userError.GetHttpStatus()
        publicMessage = net.StatusText(status)
    } else {
        // Handle unknown error as a system error
        // Please join the error with with a user error:
        // errors.Join(err, entity.UserError{HttpStatus: net.StatusUnauthorized})
        status = net.StatusInternalServerError
        publicMessage = "unhandled internal server error"
        internalMessage = err.Error()
    }

    // Set the response header to application/json
    writer.Header().Set("Content-Type", "application/json")
    // Write the status code as 500 Internal Server Error
    writer.WriteHeader(status)
    // Create a JSON response with the error publicMessage
    _ = json.NewEncoder(writer).Encode(map[string]string{"report": report, "publicMessage": publicMessage})

    // If it is a user error, we don't want to log that
    if errors.As(err, &userError) {
        return
    }
    _, err = fmt.Printf("Error report: %s, message: %s", report, internalMessage)
    if err != nil {
        log.Panicln("can't print error: " + err.Error())
    }
}

func appendApiByPath(routes []entity.Route) []entity.Route {
    for _, r := range routes {
        pattern := getApiByPathPattern(r.Pattern)
        routes = append(routes, entity.Route{
            Pattern:     pattern,
            Controller:  r.Controller,
            Middlewares: r.Middlewares,
        })
    }
    return routes
}

func getApiByPathPattern(pattern string) string {
    index := strings.Index(pattern, " ")
    // With method
    // GET /images/ to GET /conf_api/repo/service/images/
    // Consider that the space can't be further than index 7 when using method `OPTIONS `
    if index != -1 && index <= 7 {
        return fmt.Sprintf(
            "%s /%s/%s%s",
            pattern[:index],
            config.AppInfo.ApiByPathPrefix,
            config.AppInfo.Service,
            pattern[index+1:],
        )
    }

    // Without method
    // /images/ to /conf_api/repo/service/images/
    return fmt.Sprintf(
        "/%s/%s%s",
        config.AppInfo.ApiByPathPrefix,
        config.AppInfo.Service,
        pattern,
    )
}
