package controller

import (
	"strconv"

	"github.com/labbs/faker/config"
	"github.com/labbs/faker/modules/github-api/domain"
	"github.com/labbs/faker/modules/github-api/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type GithubActionsRunnerController struct {
	Logger zerolog.Logger
}

func (garc *GithubActionsRunnerController) ListRunners(c *fiber.Ctx) error {
	var runners []domain.GithubActionsRunner
	var count int

	for _, runnerType := range config.AppConfig.GithubApi.Runners {
		for i := 0; i < runnerType.Count; i++ {
			count++
			busy := internal.RandomBool()
			runners = append(runners, domain.GithubActionsRunner{
				Id:     count,
				Os:     runnerType.Os,
				Name:   runnerType.Name + "_" + strconv.Itoa(count),
				Busy:   busy,
				Status: internal.RandomStatus(busy),
				Labels: internal.GenerateRunnerLabels(runnerType),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total_count": count,
		"runners":     runners,
	})
}
