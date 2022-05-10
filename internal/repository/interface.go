package repository

import (
	"github.com/ralugr/filter-service/internal/model"
)

type Base interface {
	Store(message *model.Message) error
	GetMessages(state string) ([]model.Message, error)
}
