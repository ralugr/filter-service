package model

import "fmt"

// Message type, used al over the application for passing and processing messages
type Message struct {
	UID    string `json:"id"`
	Body   string `json:"body"`
	State  string
	Reason string
}

// String used for printing and logging
func (m Message) String() string {
	return fmt.Sprintf("Message { UID: %s, Body: %s, State: %s, Reason: %s}", m.UID, m.Body, m.State, m.Reason)
}
