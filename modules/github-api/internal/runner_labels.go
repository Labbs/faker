package internal

import (
	"github.com/labbs/faker/modules/github-api/domain"
)

func GenerateRunnerLabels(os, arch string, sh bool) []domain.GithubLabel {
	var labels []domain.GithubLabel

	switch os {
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

	switch arch {
	case "x64":
		labels = append(labels, domain.GithubLabel{
			Id:   7,
			Type: "read-only",
			Name: "X64",
		})
	}

	if sh {
		labels = append(labels, domain.GithubLabel{
			Id:   5,
			Type: "read-only",
			Name: "self-hosted",
		})
	}

	return labels
}
