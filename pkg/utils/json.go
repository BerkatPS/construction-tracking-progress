package utils

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}

}

func JSONErrorResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	JSONResponse(w, statusCode, map[string]interface{}{
		"status":  "error",
		"message": message,
	})
}
