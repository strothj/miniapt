package main

import (
	"log"
	"os"

	"github.com/strothj/miniapt/commands"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	log.SetFlags(0)
	app := cli.NewApp()
	app.Name = "miniapt"
	app.Usage = "download and extract APT packages"
	commands.ConfigureCommandLine(app)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
