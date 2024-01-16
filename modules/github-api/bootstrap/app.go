package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	initBootstrap "github.com/labbs/faker/bootstrap"
)

type Application struct {
	Fiber  *fiber.App
	Logger zerolog.Logger
}

func App(initBootstrapApp *initBootstrap.Application) Application {
	app := &Application{}
	app.Logger = InitLogger(initBootstrapApp.Logger)
	app.Fiber = initBootstrapApp.Fiber
	return *app
}
