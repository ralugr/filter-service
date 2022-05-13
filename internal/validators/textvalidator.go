package validators

import (
	"bufio"
	"fmt"
	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
	"regexp"
	"strings"
)

// TextValidator responsible for checking the syntax of the message
type TextValidator struct {
}

// NewTextValidator constructor
func NewTextValidator() *TextValidator {
	logger.Info.Println("Creating text validator")
	return &TextValidator{}
}

// Validate checks if the message starts with a heading and contains at lest one paragraph
func (tv *TextValidator) Validate(message *model.Message) error {
	logger.Info.Println("Entering text validator with message ", message)

	if tv.startsWithHeading1(message) && tv.hasParagraph(message) {
		message.State = model.Accepted
		logger.Info.Printf("Message %v accepted by Text Validator ", message)

		return nil
	}

	message.State = model.Rejected
	message.Reason = model.TextValidationFailed
	logger.Info.Printf("Message %v rejected by Text Validator ", message)

	return nil
}

// startsWith checks whether the message starts with heading 1
func (tv *TextValidator) startsWithHeading1(message *model.Message) bool {
	pattern := regexp.MustCompile(`^#[^#\n].*`)
	if pattern.MatchString(message.Body) {
		return true
	}

	logger.Warning.Printf("The message %v does not start with a heading 1", message)
	return false
}

// hasParagraph checks whether a message contains at least one paragraph (a line that starts with a letter)
func (tv *TextValidator) hasParagraph(message *model.Message) bool {
	fmt.Println(message)
	scanner := bufio.NewScanner(strings.NewReader(message.Body))
	pattern := regexp.MustCompile(`^ *[a-zA-Z]+.*$`)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if pattern.MatchString(scanner.Text()) {
			return true
		}
	}

	logger.Warning.Printf("The message %v does not contain a paragraph", message)
	return false
}
