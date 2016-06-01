package commands

import "gopkg.in/urfave/cli.v2"

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
	panic("Not Implemented")
}
