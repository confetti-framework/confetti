package status

import (
	"net/http"
	"src/internal/pkg/handler"
)

// Index returns the status of the application
func Index(response http.ResponseWriter, req *http.Request) error {
	// Create a map to store the CPU information
	data := make(map[string]any)

	// Get the number of CPUs and store in the map
	data["status"] = "active"

	return handler.ToJson(response, data, http.StatusOK)
}
