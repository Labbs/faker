package main

import (
	"github-api/api/router"
	"github-api/bootstrap"
	"github-api/config"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

var version = "development"

func main() {
	flags := bootstrap.Flags()

	app := cli.NewApp()
	app.Name = "github-api"
	app.Usage = "A simple API to mock GitHub API"
	app.Version = version

	app.Flags = flags
	app.Before = altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config"))
	app.Action = func(c *cli.Context) error {
		appBootstrap := bootstrap.NewApplication()

		router.Setup(*appBootstrap)

		appBootstrap.Logger.Info().Str("event", "server.start").Int("port", config.AppConfig.Port).Msg("Starting server")
		return appBootstrap.Fiber.Listen(":" + strconv.Itoa(config.AppConfig.Port))
	}

	config.AppConfig.Version = version
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
