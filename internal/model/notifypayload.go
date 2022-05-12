package model

import "fmt"

type NotifyPayload struct {
	Token string   `json:"token"`
	Words []string `json:"words"`
}

func (n NotifyPayload) String() string {
	return fmt.Sprintf("NotifyPayload { Token: %s, Words: %s}", n.Token, n.Words)
}
