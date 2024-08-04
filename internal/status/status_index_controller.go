package status

import (
	"net/http"
	"src/internal/pkg/handler"
)

// Index returns the status of the application
func Index(response http.ResponseWriter, request *http.Request) error {
	// Create a map to store the information
	data := map[string]any{}

	// For now, the system is active
	data["status"] = "active"

	return handler.ToJson(response, data, http.StatusOK)
}
