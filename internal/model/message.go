package model

import "fmt"

type Message struct {
	UID    string `json:"id"`
	Body   string `json:"body"`
	State  string
	Reason string
}

func (m Message) String() string {
	return fmt.Sprintf("Message { UID: %s, Body: %s, State: %s, Reason: %s}", m.UID, m.Body, m.State, m.Reason)
}
