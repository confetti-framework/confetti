package routes

import (
    "encoding/json"
    "net/http"
    "src/app/http/controllers"
)

var Api = []Route{
    newRoute("/ping", controllers.Ping),
}

func HandleApiRoute(response http.ResponseWriter, request *http.Request, controller Controller) {
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
