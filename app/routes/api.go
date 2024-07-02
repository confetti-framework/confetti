package routes

import (
    "encoding/json"
    "net/http"
    "src/app/http/controllers"
)

type controller func(w http.ResponseWriter, req *http.Request) error

var Api = func(mux *http.ServeMux) {
    // Here you can register all your routes
    apiRoute(mux, "/ping", controllers.Ping)
}

func apiRoute(mux *http.ServeMux, pattern string, handler controller) {
    mux.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
        // Here you can:
        // - call middlewares to change the request and reponse
        // - use http.NewResponseController(w) to change server options like timeouts
        err := handler(writer, request)
        if err != nil {
            apiErrorHandler(writer, err)
        }
    })
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
