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

func newRoute(pattern string, controller controller) route {
    return route{pattern: pattern, controller: controller}
}

type controller func(w http.ResponseWriter, req *http.Request) error

var Api = groupApiRoutes([]route{
    newRoute("/ping", controllers.Ping),
})

func groupApiRoutes(routes []route) func(mux *http.ServeMux) {
    return func(mux *http.ServeMux) {
        for _, route := range routes {
            mux.HandleFunc(route.pattern, func(writer http.ResponseWriter, request *http.Request) {
                // Here you can:
                // - call middlewares to change the request and reponse
                // - use http.NewResponseController(w) to change server options like timeouts
                err := route.controller(writer, request)
                if err != nil {
                    apiErrorHandler(writer, err)
                }
            })
        }
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
