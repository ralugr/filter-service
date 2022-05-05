package model

import "fmt"

type Message struct {
	UID   string `json:"id"`
	Body  string `json:"body"`
	State MsgState
}

func (m Message) String() string {
	return fmt.Sprintf("Message { UID: %s, Body: %s, State: %s}", m.UID, m.Body, m.State)
}
