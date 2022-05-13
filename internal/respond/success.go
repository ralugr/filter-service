package respond

import (
	"encoding/json"
	"github.com/ralugr/filter-service/internal/logger"
	"net/http"
)

// successResponse defines a success
type successResponse struct {
	Success  bool        `json:"success"`
	Response interface{} `json:"response"`
}

// Success used for responding to a successful request
func Success(w http.ResponseWriter, response interface{}) {
	resp := successResponse{
		Success:  true,
		Response: response,
	}

	body, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if _, err := w.Write(body); err != nil {
		logger.Warning.Printf("Could not marshall response %v", err)
	}
}
