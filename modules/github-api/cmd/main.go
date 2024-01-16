package cmd

import (
	initBootstrap "github.com/labbs/faker/bootstrap"
	"github.com/labbs/faker/modules/github-api/api/router"
	"github.com/labbs/faker/modules/github-api/bootstrap"
)

func Init(initBootstrapApp *initBootstrap.Application) {
	app := bootstrap.App(initBootstrapApp)

	router.Setup(app)
}
