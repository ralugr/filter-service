package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ralugr/filter-service/internal/logger"
)

// Config is used to store all the user config values from config.json
type Config struct {
	// Name of the database file
	DBName string `json:"DBName"`

	// The port this application runs on
	Port int `json:"Port"`

	// Token send to the language-service for authentication purposes.
	// Filter-service contains a /notify route used for receiving banned list updates.
	// Only services that have this Token are able send a notification on this route.
	Token string `json:"Token"`

	// Url for retrieving the banned words list. Used at application startup only.
	BannedWordsUrl string `json:"BannedWordsUrl"`

	// Url used to subscribe to the language-service and get the banned words list updates.
	SubscribeUrl string `json:"SubscribeUrl"`

	// The url at which the language-service send us the updated list of banned words
	NotifyUrl string `json:"NotifyUrl"`
}

// New constructor
func New(file string) (*Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		logger.Warning.Println("Could not load config file ", err)
		return nil, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		logger.Warning.Println("Could not decode config file ", err)
		return nil, err
	}
	logger.Info.Println("Created the following app config: ", config)

	return &config, nil
}

// String used for printing and logging
func (c *Config) String() string {
	return fmt.Sprintf("Config { DBName: %s, Port: %d, Token: %s, BannedWordsUrl: %s, SubscribeUrl: %s}", c.DBName, c.Port, c.Token, c.BannedWordsUrl, c.SubscribeUrl)
}
