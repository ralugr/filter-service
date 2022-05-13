package repository

import (
	"github.com/ralugr/filter-service/internal/model"
)

// Base type for all repositories
type Base interface {
	// Store used to add a new message
	Store(*model.Message) error

	// GetMessages used to retrieve messages based on the given state
	GetMessages(string) ([]model.Message, error)

	// UpdateBannedWords saves the banned words into the repository
	UpdateBannedWords(*model.BannedWords) error

	// GetBannedWords returns the banned words list
	GetBannedWords() (*model.BannedWords, error)
}
