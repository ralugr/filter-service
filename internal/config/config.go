package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ralugr/filter-service/internal/logger"
)

type Config struct {
	// Name of the database file
	DBName string `json:"DBName"`
	// The port thhis application runs on
	Port int `json:"Port"`
	// Token send to the language service for authentication purposes.
	// We exposed a notify route that updates the banned words list
	// that the language filter uses. We only want the services with
	// this token to be able to do that.
	Token string `json:"Token"`
	// Url from with we get the bad words from. This happens when application starts.
	BannedWordsUrl string `json:"BannedWordsUrl"`
	// Url of the language service that we call on this service starts. The url tells
	// the language service that we want to be notified whenever its banned words list
	// is updated
	SubscribeUrl string `json:"SubscribeUrl"`
	// The url at which the language service send us the updated list of banned words
	NotifyUrl string `json:"NotifyUrl"`
}

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

func (c *Config) String() string {
	return fmt.Sprintf("Config { DBName: %s, Port: %d, Token: %s, BannedWordsUrl: %s, SubscribeUrl: %s}", c.DBName, c.Port, c.Token, c.BannedWordsUrl, c.SubscribeUrl)
}
