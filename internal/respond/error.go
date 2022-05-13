package respond

import (
	"encoding/json"
	"github.com/ralugr/filter-service/internal/logger"
	"net/http"
)

// errorResponse defines an error type
type errorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// Error used for responding with an error status
func Error(w http.ResponseWriter, status int, error string) {
	resp := errorResponse{
		Success: false,
		Error:   error,
	}

	body, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(body); err != nil {
		logger.Warning.Printf("Could not write error response %v", err)
		return
	}
}
