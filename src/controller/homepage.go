package controller

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing/outcome"
	"lanvard/resources/views"
)

func Book(request inter.Request) inter.Response {
	return outcome.Json(views.Book("James"))
}
