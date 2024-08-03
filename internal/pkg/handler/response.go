package handler

import (
	"encoding/json"
	"net/http"
)

func ToJson(response http.ResponseWriter, data any, status int) error {
	// Convert the map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Write the JSON to the response
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	_, err = response.Write(jsonData)
	return err
}
