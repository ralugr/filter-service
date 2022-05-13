package model

import "fmt"

// NotifyPayload type, defines the model we receive on notification from the language-service
type NotifyPayload struct {
	Token string   `json:"token"`
	Words []string `json:"words"`
}

// String used for printing and logging
func (n NotifyPayload) String() string {
	return fmt.Sprintf("NotifyPayload { Token: %s, Words: %s}", n.Token, n.Words)
}
