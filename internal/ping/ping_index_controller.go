package ping

import (
	"fmt"
	"net/http"
)

// Index is an endpoint with which you can check whether your application is responding.
func Index(response http.ResponseWriter, request *http.Request) error {
	_, err := fmt.Fprintf(response, "pong")
	return err
}
