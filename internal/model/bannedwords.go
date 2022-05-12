package model

import "fmt"

const ID = 1

type BannedWords struct {
	Id    int
	Words []string
}

func NewBannedWords(words []string) *BannedWords {
	bw := BannedWords{
		Id:    ID,
		Words: words,
	}

	return &bw
}

func (bw BannedWords) String() string {
	return fmt.Sprintf("BannedWords { Id: %d, Words: %s}", bw.Id, bw.Words)
}
