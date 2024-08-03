package handler

import (
	"fmt"
	"net/http"
	"src/config"
	"strings"
)

type HttpStatusGetter interface {
	GetHttpStatus() int
}

func New(pattern string, controller Controller) Route {
	return Route{Pattern: pattern, Controller: controller}
}

func RegisterRoutes(routes []Route, routeHandler func(http.ResponseWriter, *http.Request, Route)) func(mux *http.ServeMux) {
	return func(mux *http.ServeMux) {
		for _, r := range routes {
			mux.HandleFunc(r.Pattern, func(response http.ResponseWriter, request *http.Request) {
				routeHandler(response, request, r)
			})
		}
	}
}

func HandleApiRoute(response http.ResponseWriter, request *http.Request, route Route) {
	handler := MiddlewareChain{Middlewares: route.GetMiddlewares()}

	err := handler.Handle(route.Controller)(response, request)
	if err != nil {
		apiErrorHandler(response, err)
	}
}

func AppendApiByPath(routes []Route) []Route {
	if config.AppInfo.ServiceUriPrefix == "" {
		return routes
	}
	// When using this in Confetti CMS we want to register the endoints so
	// we can call them internaly with a server prefix but without a subdomain.
	for _, r := range routes {
		pattern := getApiByPathPattern(r.Pattern)
		routes = append(routes, Route{
			Pattern:     pattern,
			Controller:  r.Controller,
			Middlewares: r.Middlewares,
		})
	}
	return routes
}

func getApiByPathPattern(pattern string) string {
	index := strings.Index(pattern, " ")
	// With method
	// GET /images/ to GET /conf_api/repo/service/images/
	// Consider that the space can't be further than index 7 when using method `OPTIONS `
	if index != -1 && index <= 7 {
		return fmt.Sprintf(
			"%s %s%s",
			pattern[:index],
			config.AppInfo.ServiceUriPrefix,
			pattern[index+1:],
		)
	}

	// Without method
	// /images/ to /conf_api/repo/service/images/
	return fmt.Sprintf(
		"%s%s",
		config.AppInfo.ServiceUriPrefix,
		pattern,
	)
}
