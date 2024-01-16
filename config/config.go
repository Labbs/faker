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
	LinuxRunners   Runner `json:"linux_runners"`
	WindowsRunners Runner `json:"windows_runners"`
	MacOSRunners   Runner `json:"macos_runners"`
}

type Runner struct {
	Count      int    `json:"count"`
	Arch       string `json:"arch"`
	SelfHosted bool   `json:"self-hosted"`
}
