package service

import (
	"bytes"
	"encoding/json"
	"github.com/ralugr/filter-service/internal/config"
	"github.com/ralugr/filter-service/internal/handlers"
	"github.com/ralugr/filter-service/internal/logger"
	"github.com/ralugr/filter-service/internal/model"
	"github.com/ralugr/filter-service/internal/processor"
	"github.com/ralugr/filter-service/internal/repository"
	"github.com/ralugr/filter-service/internal/validators"
	"io/ioutil"
	"net/http"
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

	if err = subscribe(cfg); err != nil {
		logger.Warning.Printf("Unable to subscribe to language service %v", err)
	}

	// Call the language service and get the words list. This way we make sure
	// that whenever we start this service we have the updated list of banned words.
	bannedWords, err := getBannedWords(cfg)

	if len(bannedWords) != 0 {
		bw := model.NewBannedWords(bannedWords)
		repo.UpdateBannedWords(bw)
	}

	// Initialized processor with concrete types.
	p := processor.New(
		[]validators.Base{validators.NewLinkValidator(), validators.NewImageValidator(), validators.NewTextValidator(), validators.NewLanguageValidator(repo)},
		repo,
	)
	h := handlers.New(p, cfg)

	return &Service{cfg, h, p}, nil
}

func getBannedWords(cfg *config.Config) ([]string, error) {
	var bannedWords []string
	response, err := http.Get(cfg.BannedWordsUrl)

	if err != nil {
		logger.Warning.Printf("Failed to get banned words from language service %v", err)
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Warning.Printf("Failed to read response body %v", err)
		return nil, err
	}

	err = json.Unmarshal(responseData, &bannedWords)
	if err != nil {
		logger.Warning.Printf("Failed to decode response body %v", err)
		return nil, err
	}

	return bannedWords, nil
}

func subscribe(cfg *config.Config) error {
	subscriber := model.Subscriber{
		Token: cfg.Token,
		URL:   cfg.NotifyUrl,
	}

	data, err := json.Marshal(subscriber)
	if err != nil {
		logger.Info.Printf("Could not marshall subscriber %v", subscriber)
		return err
	}

	_, err = http.Post(cfg.SubscribeUrl, "application/json", bytes.NewReader(data))

	if err != nil {
		logger.Warning.Printf("POST failed for subscription %v, error %v", subscriber, err)
		return err
	}

	logger.Info.Printf("Subscribe sent to language service")
	return nil
}
