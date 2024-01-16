package internal

import (
	"github.com/labbs/faker/config"
	"github.com/labbs/faker/modules/github-api/domain"
)

func GenerateRunnerLabels(runner config.Runner) []domain.GithubLabel {
	var labels []domain.GithubLabel

	switch runner.Os {
	case "linux":
		labels = append(labels, domain.GithubLabel{
			Id:   11,
			Type: "read-only",
			Name: "Linux",
		})
	case "macos":
		labels = append(labels, domain.GithubLabel{
			Id:   20,
			Type: "read-only",
			Name: "macOS",
		})
	case "windows":
		labels = append(labels, domain.GithubLabel{})
	}

	switch runner.Arch {
	case "x64":
		labels = append(labels, domain.GithubLabel{
			Id:   7,
			Type: "read-only",
			Name: "X64",
		})
	}

	if runner.SelfHosted {
		labels = append(labels, domain.GithubLabel{
			Id:   5,
			Type: "read-only",
			Name: "self-hosted",
		})
	}

	return labels
}
