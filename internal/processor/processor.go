package processor

import (
	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
	"github.com/ralugr/filter-service/internal/repository"
	"github.com/ralugr/filter-service/internal/validators"
	"reflect"
)

type Processor struct {
	validators []validators.Base
	repo       repository.Base
}

// New takes an sqs client and config and returns a new processor
func New(v []validators.Base, r repository.Base) *Processor {
	logger.Info.Printf("Creating a new processor with the following validators: %v and repo: %v", reflect.TypeOf(v), reflect.TypeOf(r))

	return &Processor{
		validators: v,
		repo:       r,
	}
}

func (p *Processor) FilterMessage(message model.Message) {
	logger.Info.Println("Message is being filtered: ", message)

	for _, elem := range p.validators {
		err := elem.Validate(&message)
		if err != nil {
			logger.Warning.Println("Validation failed with error: ", err) // have a package with predefined errors
		}

		if message.State == model.Rejected || message.State == model.Queued {
			logger.Info.Println("Saving message into the repo: ", message)
			err = p.repo.Store(&message)
			if err != nil {
				logger.Warning.Println("Could not save message %v into repo", message)
			}
			return
		}
	}
}

func (p *Processor) GetMessages(state string) ([]model.Message, error) {
	return p.repo.GetMessages(state)
}
