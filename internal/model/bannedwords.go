package model

import "fmt"

// ID used to indicate that the service replaces the existing list
const ID = 1

// BannedWords type, used for encapsulated the banned list
type BannedWords struct {
	Id    int
	Words []string
}

// NewBannedWords constructor
func NewBannedWords(words []string) *BannedWords {
	bw := BannedWords{
		Id:    ID,
		Words: words,
	}

	return &bw
}

// String used for printing and logging
func (bw BannedWords) String() string {
	return fmt.Sprintf("BannedWords { Id: %d, Words: %s}", bw.Id, bw.Words)
}
