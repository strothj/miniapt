package commands

import (
	"github.com/strothj/go-debrepo/debrepo"
	"github.com/strothj/miniapt/miniapt"
	"gopkg.in/urfave/cli.v2"
)

func addGlobalFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    miniapt.DataDirFlagName,
			Usage:   "set path to data directory, uses $PWD/miniapt if not set",
			EnvVars: []string{"MINIAPT_DATADIR"},
		},
		&cli.StringFlag{
			Name:    miniapt.ArchitectureFlagName,
			Usage:   "set package architecture",
			Value:   debrepo.DetectArchitecture(),
			EnvVars: []string{"MINIAPT_ARCH"},
		},
	}
}
