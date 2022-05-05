package stores

import (
	"github.com/ralugr/filter-service/pkg/model"
)

type Base interface {
	Store(message model.Message)
	GetMessages(state model.MsgState)
}
