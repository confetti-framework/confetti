package controller

import (
	"fmt"
	net "net/http"
)

// Status returns the status of the application
func Status(response net.ResponseWriter, req *net.Request) error {
	_, err := fmt.Fprintf(response, "Ok")
	return err
}
