package bootstrap

import (
	"fmt"
	"os"

	"github.com/labbs/faker/config"

	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
)

func LoadConfig(logger zerolog.Logger) {
	// Read config from file
	jsonFile, err := os.ReadFile(config.AppConfig.ConfigFile)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to read config file")
		return
	}

	// Parse config
	var data config.Config = config.Config{}

	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse config file")
		return
	}

	fmt.Println(data)

	// Set config
	config.AppConfig.GithubApi.Runners = data.GithubApi.Runners
	logger.Info().Str("event", "config.loaded").Msg("Loaded config")
}
