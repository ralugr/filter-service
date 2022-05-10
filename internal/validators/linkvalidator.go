package validators

import (
	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
	"strings"
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
	//pattern := regexp.MustCompile(`(^\[.*\]\s*\(.*\))|((?:[^!])(\[.*\]\s*\(.*\)))`)
	//links := pattern.FindAllString(message.Body, -1)

	//for _, l := range links {
	//	u := lv.match(l)
	//
	//	if u == "" {
	//		continue
	//	}
	//
	//	logger.Info.Printf("Checking link %v", u)
	//
	//	url, err := url.Parse(u)
	//
	//	if err != nil {
	//		logger.Warning.Printf("The URL %v could not be parsed %v", u, err)
	//	}
	//	if url != nil && url.IsAbs() {
	//		logger.Warning.Printf("The message %v has an external link %v", message, l)
	//		return false
	//	}
	//}

	return false // true
}

func (lv *LinkValidator) match(s string) string {
	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s, ")")
		if j >= 0 {
			return s[i+1 : j]
		}
	}
	return ""
}
