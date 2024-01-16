package bootstrap

import (
	"github.com/labbs/faker/config"

	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func GenericFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "Load configuration from `FILE`",
			Value:       "config.json",
			Destination: &config.AppConfig.ConfigFile,
		},
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "debug",
			Usage:       "Enable debug mode",
			Value:       false,
			Destination: &config.AppConfig.Debug,
		}),
	}
}

func GithubApiFlags() []cli.Flag {
	return []cli.Flag{
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "github-api.linux-runners.self-hosted",
			Usage:       "Enable self-hosted Linux runners",
			Value:       false,
			Destination: &config.AppConfig.GithubApi.LinuxRunners.SelfHosted,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "github-api.linux-runners.count",
			Usage:       "Set the number of Linux runners",
			Value:       0,
			Destination: &config.AppConfig.GithubApi.LinuxRunners.Count,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "github-api.linux-runners.arch",
			Usage:       "Set the architecture of Linux runners",
			Value:       "x64",
			Destination: &config.AppConfig.GithubApi.LinuxRunners.Arch,
		}),

		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "github-api.windows-runners.self-hosted",
			Usage:       "Enable self-hosted Windows runners",
			Value:       false,
			Destination: &config.AppConfig.GithubApi.WindowsRunners.SelfHosted,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "github-api.windows-runners.count",
			Usage:       "Set the number of Windows runners",
			Value:       0,
			Destination: &config.AppConfig.GithubApi.WindowsRunners.Count,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "github-api.windows-runners.arch",
			Usage:       "Set the architecture of Windows runners",
			Value:       "x64",
			Destination: &config.AppConfig.GithubApi.WindowsRunners.Arch,
		}),

		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "github-api.macos-runners.self-hosted",
			Usage:       "Enable self-hosted MacOS runners",
			Value:       false,
			Destination: &config.AppConfig.GithubApi.MacOSRunners.SelfHosted,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "github-api.macos-runners.count",
			Usage:       "Set the number of MacOS runners",
			Value:       0,
			Destination: &config.AppConfig.GithubApi.MacOSRunners.Count,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "github-api.macos-runners.arch",
			Usage:       "Set the architecture of MacOS runners",
			Value:       "x64",
			Destination: &config.AppConfig.GithubApi.MacOSRunners.Arch,
		}),
	}
}

func ServerFlags() []cli.Flag {
	flags := []cli.Flag{
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "enable-http-logs",
			Usage:       "Enable HTTP logs",
			Value:       false,
			Destination: &config.AppConfig.EnableHTTPLogs,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:        "pretty-logs",
			Usage:       "Enable pretty logs",
			Value:       false,
			Destination: &config.AppConfig.PrettyLogs,
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:        "port",
			Usage:       "Set the port to listen on",
			Value:       8080,
			Destination: &config.AppConfig.Port,
		}),
	}

	flags = append(flags, GenericFlags()...)
	flags = append(flags, GithubApiFlags()...)
	return flags
}
