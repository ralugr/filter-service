package model

import "fmt"

// Subscriber used to subscribe to language-service notifications
type Subscriber struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

// String used for printing and logging
func (s Subscriber) String() string {
	return fmt.Sprintf("Subscriber { Token: %s, URL: %s}", s.Token, s.URL)
}
