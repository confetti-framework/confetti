package controllers

import "lanvard/http"

type User struct {
}

func (u User) Index(request http.Request) http.Response {
	return http.Json("{\"test\":  123}")
}

func (u User) Store(request http.Request) http.Response {
	return http.NewResponse()
}

func (u User) Destroy(request http.Request) http.Response {
	return http.NewResponse()
}
