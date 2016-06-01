package commands

import (
	"github.com/strothj/miniapt/miniapt"
	"gopkg.in/urfave/cli.v2"
)

func addGlobalFlags(app *cli.App) {
	app.Flags = append(app.Flags, &cli.StringFlag{
		Name:    miniapt.DataDirFlagName,
		Usage:   "set path to data directory, uses $PWD/miniapt if not set",
		EnvVars: []string{"MINIAPT_DATADIR"},
	})
}
