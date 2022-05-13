package processor

import (
	"reflect"
	"strings"

	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
	"github.com/ralugr/filter-service/internal/repository"
	"github.com/ralugr/filter-service/internal/validators"
)

// Processor type, defined the main flow of the service
type Processor struct {
	validators []validators.Base
	repo       repository.Base
}

// New constructor
func New(v []validators.Base, r repository.Base) *Processor {
	logger.Info.Printf("Creating a new processor with the following validators: %v and repo: %v", reflect.TypeOf(v), reflect.TypeOf(r))

	return &Processor{
		validators: v,
		repo:       r,
	}
}

// FilterMessage parses all the validators and calls the validate method.
// Communicates further actions to the repository.
func (p *Processor) FilterMessage(message *model.Message) {
	logger.Info.Println("Message is being filtered: ", message)

	// Replaces <br> with \n
	newBody := strings.Replace(message.Body, "<br>", "\n", -1)
	message.Body = newBody

	for _, elem := range p.validators {
		err := elem.Validate(message)
		if err != nil {
			logger.Warning.Println("Validation failed with error: ", err) // have a package with predefined errors
		}

		if message.State == model.Rejected || message.State == model.Queued {
			logger.Info.Println("Saving message into the repo: ", message)
			err = p.repo.Store(message)
			if err != nil {
				logger.Warning.Printf("Could not save message %v into repo", message)
			}
			return
		}
	}
}

// GetMessages redirects call to the repository for retrieving rejected or queued messages
func (p *Processor) GetMessages(state string) ([]model.Message, error) {
	return p.repo.GetMessages(state)
}

// UpdateBannedWords saves the banned list to the repository
func (p *Processor) UpdateBannedWords(bw *model.BannedWords) error {
	return p.repo.UpdateBannedWords(bw)
}
