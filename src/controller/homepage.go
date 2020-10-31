package controller

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/routing/outcome"
	"lanvard/app/report"
)

func Homepage(request inter.Request) inter.Response {
	return outcome.Html(report.PageNotFoundError)
}
