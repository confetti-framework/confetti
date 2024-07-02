package controllers

import (
    "fmt"
    "net/http"
)

// Ping is an endpoint with which you can check whether your application is responding.
func Ping(w http.ResponseWriter, req *http.Request) error {
    _, err := fmt.Fprintf(w, "pong")
    return err
}
