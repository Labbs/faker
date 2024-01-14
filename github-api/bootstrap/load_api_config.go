package bootstrap

import (
	"github-api/config"
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/yaml.v2"
)

func LoadApiConfig(logger zerolog.Logger) {
	// Read config from file
	yamlFile, err := os.ReadFile(config.AppConfig.ConfigFile)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to read config file")
		return
	}

	// Parse config
	var data config.Config = config.Config{}

	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse config file")
		return
	}

	// Set config
	config.AppConfig.Runners = data.Runners
	logger.Info().Str("event", "config.loaded").Msg("Loaded config")
}
