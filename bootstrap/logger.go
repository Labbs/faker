package bootstrap

import (
	"os"

	"github.com/rs/zerolog"

	"github.com/labbs/faker/config"
)

func InitLogger() zerolog.Logger {
	host, _ := os.Hostname()
	logger := zerolog.New(os.Stderr).With().
		Caller().
		Timestamp().
		Str("host", host).
		Str("version", config.AppConfig.Version).
		Logger()

	if config.AppConfig.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if config.AppConfig.PrettyLogs {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return logger
}
