package commands

import "gopkg.in/urfave/cli.v2"

// ConfigureCommandLine adds commands to the command line application.
func ConfigureCommandLine(app *cli.App) {
	addGlobalFlags(app)

	addSourceCommand(app)
	addKeyCommand(app)
}
