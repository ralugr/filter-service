package validators

import (
	"net/url"
	"regexp"

	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
)

// LinkValidator responsible for checking for links and validating them
type LinkValidator struct {
}

func NewLinkValidator() *LinkValidator {
	logger.Info.Println("Creating link validator")
	return &LinkValidator{}
}

func (lv *LinkValidator) Validate(message *model.Message) error {
	logger.Info.Println("Entering link validator with message ", message)
	if lv.hasExternalLink(message) {
		message.State = model.Rejected
		message.Reason = model.LinkValidationFailed
	} else {
		message.State = model.Accepted
	}

	return nil
}

// hasExternalLink checks if the message has any valid URL
func (lv *LinkValidator) hasExternalLink(message *model.Message) bool {
	pattern := regexp.MustCompile(`\[[^\[\]]*\]\s*\("((?:\bhttps?:\/\/)?[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|])"\)`)
	links := pattern.FindAllStringSubmatch(message.Body, -1)

	for _, match := range links {
		fullMatch := match[0]
		matchedUrl := match[1]

		logger.Info.Printf("Checking link %v", matchedUrl)

		url, err := url.Parse(matchedUrl)

		if err != nil {
			logger.Warning.Printf("The URL %v could not be parsed %v", matchedUrl, err)
		}
		if url != nil && url.IsAbs() {
			logger.Warning.Printf("The message %v has an external link %v", message, fullMatch)
			return true
		}
	}

	return false
}
