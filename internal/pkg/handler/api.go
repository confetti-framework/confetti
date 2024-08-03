package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"src/config"
	"strings"
	"time"
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

// apiErrorHandler converts the error to a reponse. Feel free to modify this to your needs.
func apiErrorHandler(writer http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	publicMessage := ""
	internalMessage := ""
	report := fmt.Sprintf("%d", time.Now().UnixMilli())
	var systemError SystemError
	var userError UserError
	if errors.As(err, &systemError) {
		// Handle system error
		status = systemError.GetHttpStatus()
		publicMessage = "internal server error"
		internalMessage = err.Error()
	} else if errors.As(err, &userError) {
		// Handle user error
		status = userError.GetHttpStatus()
		publicMessage = userError.Message
	} else {
		// Handle unknown error as a system error
		// Please join the error with with a user error:
		// errors.Join(err, handler.UserError{HttpStatus: http.StatusUnauthorized})
		status = http.StatusInternalServerError
		publicMessage = "unhandled internal server error"
		internalMessage = err.Error()
	}

	// Set the response header to application/json
	writer.Header().Set("Content-Type", "application/json")
	// Write the status code as 500 Internal Server Error
	writer.WriteHeader(status)
	// Create a JSON response with the error publicMessage
	_ = json.NewEncoder(writer).Encode(map[string]string{
		"report":  report,
		"message": publicMessage,
	})

	// If it is a user error, we don't want to log that
	if errors.As(err, &userError) {
		return
	}
	_, err = fmt.Printf("Error report: %s, message: %s", report, internalMessage)
	if err != nil {
		log.Panicln("can't print error: " + err.Error())
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
