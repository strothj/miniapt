package commands

import (
	"fmt"
	"strings"

	"github.com/strothj/go-debrepo/debrepo"
	"gopkg.in/urfave/cli.v2"
)

func addListArchitecturesCommand(app *cli.App) {
	listArchsCommand := &cli.Command{
		Name:   "list-architectures",
		Usage:  "list architecture types",
		Action: doListArchitectures,
	}
	app.Commands = append(app.Commands, listArchsCommand)
}

func doListArchitectures(ctx *cli.Context) error {
	archs := debrepo.ListArchitectures()
	archsOut := strings.Join(archs, "\n")
	fmt.Print(archsOut)
	return nil
}
