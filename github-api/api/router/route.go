package router

import "github-api/bootstrap"

func Setup(app bootstrap.Application) {
	NewGithubActionsRunnerRouter(app)
}
