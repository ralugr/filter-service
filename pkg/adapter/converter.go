package adapter

import (
	"encoding/json"
	"github.com/ralugr/filter-service/pkg/logger"
	"github.com/ralugr/filter-service/pkg/model"
)

func ConvertJsonToMessage(jsonMsg string) model.Message {
	var message model.Message
	err := json.Unmarshal([]byte(jsonMsg), &message)
	if err != nil {
		logger.Warning.Println("Could not convert json string to message type: ", err)
	}
	logger.Info.Println("Message converted: ", message)
	return message
}
