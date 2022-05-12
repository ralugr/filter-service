package validators

import (
	"strings"

	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
	"github.com/ralugr/filter-service/internal/repository"
)

//LanguageValidator responsible for checking the language based on a banned words list
type LanguageValidator struct {
	repo repository.Base
}

func NewLanguageValidator(repo repository.Base) *LanguageValidator {
	logger.Info.Println("Creating language validator")
	return &LanguageValidator{repo: repo}
}

func (lv *LanguageValidator) Validate(message *model.Message) error {
	bannedWords, err := lv.repo.GetBannedWords()

	// If we get an error while fetching the words, consider we have no bad words, making the message state accepted
	if err != nil {
		message.State = model.Accepted
		return nil
	}

	if lv.hasBannedWords(message, bannedWords) {
		message.State = model.Rejected
		message.Reason = model.LanguageValidationFailed
	} else {
		message.State = model.Accepted
	}

	return nil
}

func (lv *LanguageValidator) hasBannedWords(message *model.Message, bannedWords *model.BannedWords) bool {
	if bannedWords == nil {
		return false
	}

	for _, bannedWord := range bannedWords.Words {
		lowerCaseBody := strings.ToLower(message.Body)
		lowerCaseBannedWord := strings.ToLower(bannedWord)
		if strings.Contains(lowerCaseBody, lowerCaseBannedWord) {
			return true
		}
	}

	return false
}
