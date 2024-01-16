package main

import (
	"log"
	"os"
	"strconv"

	"github.com/labbs/faker/bootstrap"
	"github.com/labbs/faker/config"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"

	githubApiModule "github.com/labbs/faker/modules/github-api/cmd"
)

var version = "development"

func main() {
	flags := bootstrap.ServerFlags()

	app := cli.NewApp()
	app.Name = "faker"
	app.Usage = "It's an app for mocking APIs"
	app.Version = version

	app.Flags = flags
	app.Before = altsrc.InitInputSourceWithContext(flags, altsrc.NewJSONSourceFromFlagFunc("config"))
	app.Action = func(c *cli.Context) error {
		appBootstrap := bootstrap.NewApplication()

		appBootstrap.Logger.Debug().Str("event", "app.start").Interface("config", config.AppConfig).Send()

		githubApiModule.Init(appBootstrap)

		appBootstrap.Logger.Info().Str("event", "server.start").Int("port", config.AppConfig.Port).Msg("Starting server")
		return appBootstrap.Fiber.Listen(":" + strconv.Itoa(config.AppConfig.Port))
	}

	config.AppConfig.Version = version
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
