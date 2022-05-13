package validators

import (
	"regexp"

	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
)

// ImageValidator responsible for checking for images and adding the appropriate State and Reason
type ImageValidator struct {
}

// NewImageValidator constuctor
func NewImageValidator() *ImageValidator {
	logger.Info.Println("Creating image validator")
	return &ImageValidator{}
}

// Validate check if the existing images are valid or not.There are 4 cases:
// Case 1: All images are tagged as approved: Sets message state to Approved
// Case 2: At lest one images has the rejected tag: Sets message state to Rejected
// Case 3: The images contain no approve or rejected tags: Sets message state to Queued
// Case 4: The message does not contain any images: Sets message state to Approved
func (iv *ImageValidator) Validate(message *model.Message) error {
	logger.Info.Println("Entering image validator with message ", message)

	if !iv.hasImages(message) {
		message.State = model.Accepted
		logger.Info.Printf("The message %v has no images, exiting ImageValidator ", message)

		return nil
	}

	hasRejectedTags := iv.hasTag(message, model.Rejected)
	hasAcceptedTags := iv.hasTag(message, model.Accepted)

	if hasRejectedTags {
		message.State = model.Rejected
		message.Reason = model.ImageValidationFailed
		logger.Info.Printf("Message %v rejected by image validator with reason %v", message, message.Reason)
		return nil
	}

	if hasAcceptedTags {
		message.State = model.Accepted
		logger.Info.Printf("Message %v accepted by image validator", message)
		return nil
	}

	message.State = model.Queued
	message.Reason = model.ManualValidationNeeded
	logger.Info.Printf("Message %v queued by image validator with reason %v", message, message.Reason)

	return nil
}

// hasImages checks if the message contains any images having the following syntax:
//![alt text for screen readers](/path/to/image.png "Text to show on mouseover")
func (iv *ImageValidator) hasImages(message *model.Message) bool {
	body := message.Body

	pattern := regexp.MustCompile(`!\[.*\]\(.*\)`)
	images := pattern.FindAllString(body, -1)

	return images != nil
}

// hasTag checks if the message has any tags having the following syntax, case insensitive:
// <!-- state: Rejected --> <!-- state: Approved -->
func (iv *ImageValidator) hasTag(message *model.Message, state string) bool {

	tags := iv.getTags(message, state)

	return tags != nil
}

// getTags returns all the tags found in the message
func (iv *ImageValidator) getTags(message *model.Message, state string) []string {
	body := message.Body

	pattern := regexp.MustCompile(`(?i)<!--\s*state\s*:\s*(` + state + `)\s*-->`)
	foundTags := pattern.FindAllString(body, -1)

	return foundTags
}
