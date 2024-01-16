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

	if config.AppConfig.GithubApi.LinuxRunners.Count > 0 {
		for i := 0; i < config.AppConfig.GithubApi.LinuxRunners.Count; i++ {
			count++
			busy := internal.RandomBool()
			runners = append(runners, domain.GithubActionsRunner{
				Id:     count,
				Os:     "Linux",
				Name:   "linux_" + strconv.Itoa(count),
				Busy:   busy,
				Status: internal.RandomStatus(busy),
				Labels: internal.GenerateRunnerLabels("linux",
					config.AppConfig.GithubApi.LinuxRunners.Arch,
					config.AppConfig.GithubApi.LinuxRunners.SelfHosted),
			})
		}
	}

	if config.AppConfig.GithubApi.WindowsRunners.Count > 0 {
		for i := 0; i < config.AppConfig.GithubApi.WindowsRunners.Count; i++ {
			count++
			busy := internal.RandomBool()
			runners = append(runners, domain.GithubActionsRunner{
				Id:     count,
				Os:     "Windows",
				Name:   "windows_" + strconv.Itoa(count),
				Busy:   busy,
				Status: internal.RandomStatus(busy),
				Labels: internal.GenerateRunnerLabels("windows",
					config.AppConfig.GithubApi.WindowsRunners.Arch,
					config.AppConfig.GithubApi.WindowsRunners.SelfHosted),
			})
		}
	}

	if config.AppConfig.GithubApi.MacOSRunners.Count > 0 {
		for i := 0; i < config.AppConfig.GithubApi.MacOSRunners.Count; i++ {
			count++
			busy := internal.RandomBool()
			runners = append(runners, domain.GithubActionsRunner{
				Id:     count,
				Os:     "MacOS",
				Name:   "macos_" + strconv.Itoa(count),
				Busy:   busy,
				Status: internal.RandomStatus(busy),
				Labels: internal.GenerateRunnerLabels("macos",
					config.AppConfig.GithubApi.MacOSRunners.Arch,
					config.AppConfig.GithubApi.MacOSRunners.SelfHosted),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"total_count": count,
		"runners":     runners,
	})
}
