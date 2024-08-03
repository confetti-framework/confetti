package handler

import (
	"github.com/matryer/is"
	"src/config"
	"testing"
)

func Test_append_route_by_path_returns_original_pattern(t *testing.T) {
	// Given
	routes := []Route{{Pattern: "GET /images/"}}

	// When
	routes = AppendApiByPath(routes)

	// Then
	i := is.New(t)
	i.Equal(routes[0].Pattern, "GET /images/")
}

func Test_append_route_by_path_returns_new_pattern(t *testing.T) {
	// Given
	routes := []Route{{Pattern: "GET /images/"}}
	config.AppInfo.ServiceUriPrefix = "/conf_api/confetti-cms/media"

	// When
	routes = AppendApiByPath(routes)

	// Then
	i := is.New(t)
	i.Equal(routes[1].Pattern, "GET /conf_api/confetti-cms/media/images/")
}

func Test_append_route_by_path_without_method(t *testing.T) {
	// Given
	routes := []Route{{Pattern: "/images/"}}
	config.AppInfo.ServiceUriPrefix = "/conf_api/confetti-cms/media"

	// When
	routes = AppendApiByPath(routes)

	// Then
	i := is.New(t)
	i.Equal(routes[1].Pattern, "/conf_api/confetti-cms/media/images/")
}

func Test_append_route_by_path_without_method_with_space_on_the_end(t *testing.T) {
	// Given
	routes := []Route{{Pattern: "/images/1 2"}}
	config.AppInfo.ServiceUriPrefix = "/conf_api/confetti-cms/media"

	// When
	routes = AppendApiByPath(routes)

	// Then
	i := is.New(t)
	i.Equal(routes[1].Pattern, "/conf_api/confetti-cms/media/images/1 2")
}
