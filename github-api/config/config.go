package config

var AppConfig Config = Config{}

type Config struct {
	ConfigFile     string `yaml:"config_file"`
	Debug          bool   `yaml:"debug"`
	PrettyLogs     bool   `yaml:"pretty_logs"`
	Port           int    `yaml:"port"`
	Version        string `yaml:"version"`
	EnableHTTPLogs bool   `yaml:"enable_http_logs"`

	Runners []RunnerConfig `yaml:"runners"`
}

type RunnerConfig struct {
	Os         string `yaml:"os"`
	Count      int    `yaml:"count"`
	Name       string `yaml:"name"`
	Arch       string `yaml:"arch"`
	SelfHosted bool   `yaml:"self_hosted"`
}
