package adapter

import (
	"database/sql"
	"encoding/json"

	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
)

// ConvertByteArrayToMessage converts a byte array to a message type
func ConvertByteArrayToMessage(bytes []byte) (model.Message, error) {
	var message model.Message
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		logger.Warning.Println("Could not convert json string to message type: ", err)
		return message, err
	}
	logger.Info.Println("Message converted: ", message)
	return message, nil
}

// ConvertByteArrayToNotifyPayload converts a byte array to a payload type
func ConvertByteArrayToNotifyPayload(bytes []byte) (model.NotifyPayload, error) {
	var message model.NotifyPayload
	err := json.Unmarshal(bytes, &message)
	if err != nil {
		logger.Warning.Println("Could not convert json string to message type: ", err)
		return message, err
	}
	logger.Info.Println("Message converted: ", message)
	return message, nil
}

// ConvertRowsToMessages converts rows to message type
func ConvertRowsToMessages(rows *sql.Rows) ([]model.Message, error) {
	var messages []model.Message
	// ID given at insert, but not used in message construction
	var id int

	for rows.Next() {
		var m model.Message
		if err := rows.Scan(&id, &m.UID, &m.Body, &m.State, &m.Reason); err != nil {
			return messages, err
		}
		messages = append(messages, m)
	}

	rows.Close()
	return messages, nil
}
