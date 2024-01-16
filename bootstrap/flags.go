package bootstrap

import (
	"github.com/labbs/faker/config"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func GenericFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "Load configuration from `FILE`",
			Value:       "config.json",
			Destination: &config.AppConfig.ConfigFile,
		},
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "debug",
			Usage:       "Enable debug mode",
			Value:       false,
			Destination: &config.AppConfig.Debug,
		}),
	}
}

func ServerFlags() []cli.Flag {
	flags := []cli.Flag{
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "enable-http-logs",
			Usage:       "Enable HTTP logs",
			Value:       false,
			Destination: &config.AppConfig.EnableHTTPLogs,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "pretty-logs",
			Usage:       "Enable pretty logs",
			Value:       false,
			Destination: &config.AppConfig.PrettyLogs,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "port",
			Usage:       "Set the port to listen on",
			Value:       8080,
			Destination: &config.AppConfig.Port,
		}),
	}

	flags = append(flags, GenericFlags()...)

	return flags
}
