package config

var AppConfig Config = Config{}

type Config struct {
	ConfigFile     string `json:"-"`
	Version        string `json:"-"` // This is set at build time
	Debug          bool   `json:"debug"`
	PrettyLogs     bool   `json:"pretty-logs"`
	Port           int    `json:"port"`
	EnableHTTPLogs bool   `json:"enable-http-logs"`

	GithubApi GithubApi `json:"github_api"`
}

type GithubApi struct {
	Runners []Runner `json:"runners"`
}

type Runner struct {
	Os         string `json:"os"`
	Count      int    `json:"count"`
	Name       string `json:"name"`
	Arch       string `json:"arch"`
	SelfHosted bool   `json:"self-hosted"`
}
