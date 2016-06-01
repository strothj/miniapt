package main

import (
	"os"

	"github.com/strothj/miniapt/commands"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "miniapt"
	app.Usage = "download and extract APT packages"
	commands.ConfigureCommandLine(app)
	app.Run(os.Args)
}
