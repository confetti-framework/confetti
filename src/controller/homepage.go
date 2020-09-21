package controller

import (
	"fmt"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing/outcome"
)

func Homepage(request inter.Request) inter.Response {
	source := request.Source()
	err := source.ParseForm()
	if err != nil {
		panic(err)
	}
	fmt.Println(source.FormFile("the_key"))
	return outcome.Html("ok")
}
