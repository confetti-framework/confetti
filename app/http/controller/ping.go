package controller

import (
    "fmt"
    net "net/http"
)

// Ping is an endpoint with which you can check whether your application is responding.
func Ping(response net.ResponseWriter, request *net.Request) error {
    _, err := fmt.Fprintf(response, "pong")
    return err
}
