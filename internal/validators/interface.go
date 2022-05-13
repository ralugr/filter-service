package validators

import (
	"github.com/ralugr/filter-service/internal/model"
)

// Base interface for all validators
type Base interface {
	// Validate method used for submitting a message for validation.
	// The message object will be altered with the appropriate State and Reason
	Validate(*model.Message) error
}
