package controllers

import (
    "net/http"
)

// Status returns the status of the application
func Status(response http.ResponseWriter, req *http.Request) error {
    println("the line")
    return nil
}
