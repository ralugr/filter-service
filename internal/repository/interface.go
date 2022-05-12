package repository

import (
	"github.com/ralugr/filter-service/internal/model"
)

type Base interface {
	Store(*model.Message) error
	GetMessages(string) ([]model.Message, error)
	UpdateBannedWords(*model.BannedWords) error
	GetBannedWords() (*model.BannedWords, error)
}
