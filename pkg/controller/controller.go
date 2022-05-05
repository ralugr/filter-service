package controller

import (
	"github.com/ralugr/filter-service/pkg/adapter"
	"github.com/ralugr/filter-service/pkg/logger"
	"github.com/ralugr/filter-service/pkg/model"
)

func FilterMessage(jsonMsg string) {
	message := adapter.ConvertJsonToMessage(jsonMsg)
	logger.Info.Println("Message is being filtered: ", message)

	for _, elem := range Validators {
		err := elem.Validate(&message)
		if err != nil {
			logger.Warning.Println("Validation failed with error: ", err)
		}

		if message.State == model.Rejected || message.State == model.Queued {
			logger.Info.Println("Saving message into the store: ", message)
			Store.Store(message)
			return
		}
	}
}

func getRejectedMessages() {

}
