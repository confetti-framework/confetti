package controllers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/outcome"
)

// Homepage contains the code responsible for generating the home page.
func Homepage(request inter.Request) inter.Response {
	db := request.App().Db()
	databases := db.Query("SHOW databases;")
	if databases.Empty() {

		return outcome.Html("Database empty")
	}
	return outcome.Html("Databases: "+databases.First().String())
}
