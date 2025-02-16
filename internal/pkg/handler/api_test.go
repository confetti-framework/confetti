package handler

import (
	"src/config"
	"testing"
)

func Test_append_route_by_path_returns_original_pattern(t *testing.T) {
	// Given
	routes := []Route{{Pattern: "GET /images/"}}

	// When
	routes = AppendApiByPath(routes)

	// Then
	if routes[0].Pattern != "GET /images/" {
		t.Errorf("expected %v, got %v", "GET /images/", routes[0].Pattern)
	}
}

func Test_append_route_by_path_returns_new_pattern(t *testing.T) {
	// Given
	routes := []Route{{Pattern: "GET /images/"}}
	config.AppInfo.ServiceUriPrefix = "/conf_api/confetti-cms/media"

	// When
	routes = AppendApiByPath(routes)

	// Then
	if len(routes) < 2 || routes[1].Pattern != "GET /conf_api/confetti-cms/media/images/" {
		t.Errorf("expected %v, got %v", "GET /conf_api/confetti-cms/media/images/", routes[1].Pattern)
	}
}

func Test_append_route_by_path_without_method(t *testing.T) {
	// Given
	routes := []Route{{Pattern: "/images/"}}
	config.AppInfo.ServiceUriPrefix = "/conf_api/confetti-cms/media"

	// When
	routes = AppendApiByPath(routes)

	// Then
	if len(routes) < 2 || routes[1].Pattern != "/conf_api/confetti-cms/media/images/" {
		t.Errorf("expected %v, got %v", "/conf_api/confetti-cms/media/images/", routes[1].Pattern)
	}
}

func Test_append_route_by_path_without_method_with_space_on_the_end(t *testing.T) {
	// Given
	routes := []Route{{Pattern: "/images/1 2"}}
	config.AppInfo.ServiceUriPrefix = "/conf_api/confetti-cms/media"

	// When
	routes = AppendApiByPath(routes)

	// Then
	if len(routes) < 2 || routes[1].Pattern != "/conf_api/confetti-cms/media/images/1 2" {
		t.Errorf("expected %v, got %v", "/conf_api/confetti-cms/media/images/1 2", routes[1].Pattern)
	}
}
