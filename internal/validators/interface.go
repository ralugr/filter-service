package validators

import (
	"github.com/ralugr/filter-service/internal/model"
)

type Base interface {
	Validate(message *model.Message) error
}
