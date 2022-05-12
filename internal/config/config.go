package config

import (
	"encoding/json"
	"fmt"
	"github.com/ralugr/filter-service/internal/logger"
	"os"
)

type Config struct {
	DBName         string `json:"DBName"`
	Port           int    `json:"Port"`
	Token          string `json:"Token"`
	BannedWordsUrl string `json:"BannedWordsUrl"`
	SubscribeUrl   string `json:"SubscribeUrl"`
	NotifyUrl      string `json:"NotifyUrl"`
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
