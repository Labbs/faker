package router

import (
	"github.com/labbs/faker/modules/github-api/api/controller"
	"github.com/labbs/faker/modules/github-api/bootstrap"
)

func NewGithubActionsRunnerRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Setting up Github Actions Runner routes")
	_controller := &controller.GithubActionsRunnerController{Logger: app.Logger}

	app.Fiber.Get("/orgs/:org/actions/runners", _controller.ListRunners)
}
