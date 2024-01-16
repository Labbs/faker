package router

import "github.com/labbs/faker/modules/github-api/bootstrap"

func Setup(app bootstrap.Application) {
	NewGithubActionsRunnerRouter(app)
}
