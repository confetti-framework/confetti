package controllers

import (
    "net/http"
)

// UsersMeShow returns the current user
func UsersMeShow(response http.ResponseWriter, req *http.Request) error {
	println("the line")
    return nil
}
