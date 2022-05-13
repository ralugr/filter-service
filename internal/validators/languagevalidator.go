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

// NewLanguageValidator constructor
func NewLanguageValidator(repo repository.Base) *LanguageValidator {
	logger.Info.Println("Creating language validator")
	return &LanguageValidator{repo: repo}
}

// Validate checks if the message contains any word from the banned list
func (lv *LanguageValidator) Validate(message *model.Message) error {
	logger.Info.Println("Entering image validator with message ", message)

	// Retrieves the banned words list from the internal db
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

// hasBannedWords checks if the given message contains an words from the given list
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
