package service

import (
	"github.com/ralugr/filter-service/internal/config"
	"github.com/ralugr/filter-service/internal/handlers"
	"github.com/ralugr/filter-service/internal/processor"
	"github.com/ralugr/filter-service/internal/repository"
	"github.com/ralugr/filter-service/internal/validators"
)

// Service holds the dependencies and config required for the HTTP service
type Service struct {
	// Cfg holds user configurable arguments
	Cfg *config.Config
	//Handlers handlers for HTTP requests
	Handlers *handlers.Handlers
	// Processor processes messages
	Processor *processor.Processor
}

// New creates a new service
func New(configFile string) (*Service, error) {
	cfg, err := config.New(configFile)
	if err != nil {
		return nil, err
	}

	repo, err := repository.NewSQLiteDB(cfg) // Creates concrete type
	if err != nil {
		return nil, err
	}

	// Initialized processor with concrete types.
	p := processor.New(
		[]validators.Base{validators.NewLinkValidator(), validators.NewImageValidator(), validators.NewTextValidator()},
		repo,
	)
	h := handlers.New(*p)

	return &Service{cfg, h, p}, nil
}
