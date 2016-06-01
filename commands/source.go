package commands

import (
	"errors"

	"github.com/strothj/debrepo"
	"github.com/strothj/miniapt/miniapt"
	"gopkg.in/urfave/cli.v2"
)

func addSourceCommand(app *cli.App) {
	sourceCommand := &cli.Command{
		Name:  "source",
		Usage: "manage APT repositories",
	}
	app.Commands = append(app.Commands, sourceCommand)
	addSourceAddCommand(sourceCommand)
}

func addSourceAddCommand(parent *cli.Command) {
	parent.Subcommands = append(parent.Subcommands, &cli.Command{
		Name:   "add",
		Usage:  "add APT repository",
		Action: doSourceAdd,
	})
}

func doSourceAdd(ctx *cli.Context) error {
	env, err := miniapt.EnvironmentFromContext(ctx)
	if err != nil {
		return err
	}
	if ctx.Args().Len() != 1 {
		return errors.New("Expected single argument with repository, perhaps you forgot to surround with quotes")
	}
	s, err := debrepo.ParseSource(ctx.Args().First())
	if err != nil {
		return err
	}
	sourceList := env.LoadSources()
	sourceList = append(sourceList, s)
	return env.SaveSources(sourceList)
}
