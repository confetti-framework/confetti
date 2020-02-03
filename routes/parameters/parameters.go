package parameters

import (
	"fmt"
	"github.com/lanvard/contract/rules"
	"github.com/lanvard/foundation"
	"lanvard/app/model"
	"net/url"
)

type Params struct {
	app         *foundation.Application
	urlParams   url.Values
	queryParams url.Values
}

func NewParameters(app *foundation.Application, vars map[string]string, queryParams url.Values) Params {
	urlParams := url.Values{}

	for key, value := range vars {
		urlParams.Add(key, value)
	}

	return Params{app: app, urlParams: urlParams, queryParams: queryParams}
}

func (p Params) GetUser(rules ...rules.Rule) (model.User, []error) {
	userId := p.urlParams.Get("user_id")
	fmt.Println(userId)

	// Todo; move to one validation place
	var errors []error

	for _, rule := range rules {
		err := rule.Passes(userId)

		if err != nil {
			errors = append(errors, err)
		}
	}

	var user model.User

	if errors != nil {
		return user, errors
	}

	return model.Query(), nil
}
