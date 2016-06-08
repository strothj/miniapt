package commands

import (
	"errors"

	"github.com/strothj/go-debrepo/debrepo"
	"github.com/strothj/miniapt/miniapt"
	"gopkg.in/urfave/cli.v2"
)

func addRepositoryCommand(app *cli.App) {
	repositoryCommand := &cli.Command{
		Name:  "repository",
		Usage: "manage APT repositories",
	}
	app.Commands = append(app.Commands, repositoryCommand)
	addRepositoryAddCommand(repositoryCommand)
}

func addRepositoryAddCommand(parent *cli.Command) {
	parent.Subcommands = append(parent.Subcommands, &cli.Command{
		Name:  "add",
		Usage: "add APT repository",
		Description: `Add an APT repository:
	miniapt repository add "deb http://ftp.debian.org/debian squeeze main contrib non-free"`,
		Action: doRepositoryAdd,
	})
}

func doRepositoryAdd(ctx *cli.Context) error {
	env, err := miniapt.EnvironmentFromContext(ctx)
	if err != nil {
		return err
	}
	if ctx.Args().Len() != 1 {
		return errors.New("Expected single argument with repository, perhaps you forgot to surround with quotes")
	}
	s, err := debrepo.ParseRepository(ctx.Args().First())
	if err != nil {
		return err
	}
	repoList := env.LoadRepositories()
	repoList = append(repoList, s)
	return env.SaveRepositories(repoList)
}
