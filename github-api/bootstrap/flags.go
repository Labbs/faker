package bootstrap

import (
	"github-api/config"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "Load configuration from `FILE`",
			Value:       "config.yaml",
			Destination: &config.AppConfig.ConfigFile,
		},
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "debug",
			Usage:       "Enable debug mode",
			Value:       false,
			Destination: &config.AppConfig.Debug,
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
}
