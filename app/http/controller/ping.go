package controllers

import (
    "fmt"
    "net/http"
)

// Ping is an endpoint with which you can check whether your application is responding.
func Ping(response http.ResponseWriter, req *http.Request) error {
    _, err := fmt.Fprintf(response, "pong")
    return err
}
