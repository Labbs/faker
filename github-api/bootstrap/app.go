package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Application struct {
	Fiber  *fiber.App
	Logger zerolog.Logger
}

func NewApplication() *Application {
	app := &Application{}
	app.Logger = InitLogger()
	app.Fiber = InitFiber(app.Logger)
	LoadApiConfig(app.Logger)
	return app
}
