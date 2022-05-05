package validators

import (
	"github.com/ralugr/filter-service/pkg/model"
)

type Base interface {
	Validate(message *model.Message) error
}
