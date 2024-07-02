package routes

import (
    "encoding/json"
    "net/http"
    "src/app/http/controllers"
)

type route struct {
    pattern    string
    controller controller
}

func newApiRoute(pattern string, controller controller) route {
    return route{pattern: pattern, controller: controller}
}

type controller func(w http.ResponseWriter, req *http.Request) error

var Api = GroupApiRoutes([]route{
    newApiRoute("/ping", controllers.Ping),
})

func GroupApiRoutes(routes []route) func(mux *http.ServeMux) {
    return func(mux *http.ServeMux) {
        for _, route := range routes {
            mux.HandleFunc(route.pattern, func(response http.ResponseWriter, request *http.Request) {
                HandleApiRoute(response, request, route.controller)
            })
        }
    }
}

func HandleApiRoute(response http.ResponseWriter, request *http.Request, controller controller) {
    // Here you can:
    // - call middlewares to change the request and reponse
    // - use http.NewResponseController(w) to change server options like timeouts
    err := controller(response, request)
    if err != nil {
        apiErrorHandler(response, err)
    }
}

// apiErrorHandler converts the error to a reponse. Feel free to modify this to your needs.
func apiErrorHandler(writer http.ResponseWriter, err error) {
    // Set the response header to application/json
    writer.Header().Set("Content-Type", "application/json")
    // Write the status code as 500 Internal Server Error
    writer.WriteHeader(http.StatusInternalServerError)
    // Create a JSON response with the error message
    _ = json.NewEncoder(writer).Encode(map[string]string{"message": err.Error()})
}
