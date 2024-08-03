package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetRoute(routes []Route, pattern string) Route {
	for _, route := range routes {
		if route.Pattern == pattern {
			return route
		}
	}
	panic(`Route with pattern ` + pattern + ` not found`)
}

func GetBody(result *http.Response) string {
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}

// GetField extracts a specific field from the JSON response body.
func GetField(body string, field string) any {
	var data map[string]any
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	value, exists := data[field]
	if !exists {
		panic(fmt.Sprintf("Field `%s` not found in response body", field))
	}

	return value
}
