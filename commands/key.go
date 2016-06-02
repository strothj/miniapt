package commands

import (
	"github.com/strothj/hkp"
	"github.com/strothj/miniapt/miniapt"
	"golang.org/x/net/context"
	"gopkg.in/urfave/cli.v2"
)

func addKeyCommand(app *cli.App) {
	appCommand := &cli.Command{
		Name:  "key",
		Usage: "manage PGP keys",
	}
	app.Commands = append(app.Commands, appCommand)
	addKeyAdvCommand(appCommand)
}

func addKeyAdvCommand(parent *cli.Command) {
	parent.Subcommands = append(parent.Subcommands, &cli.Command{
		Name:  "adv",
		Usage: "add key from keyserver",
		Description: `Add a key from a keyserver:
    miniapt key adv --keyserver keyserver.ubuntu.com --recv-keys 6E80C6B7`,
		Action: doKeyAdv,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "keyserver",
				Value: "keyserver.ubuntu.com",
				Usage: "keyserver to retrieve keys from",
			},
			&cli.StringFlag{
				Name:  "recv-keys",
				Usage: "key ID matching keys to retrieve",
			},
		},
	})
}

func doKeyAdv(ctx *cli.Context) error {
	env, err := miniapt.EnvironmentFromContext(ctx)
	if err != nil {
		return err
	}
	keyserver, err := hkp.ParseKeyserver(ctx.String("keyserver"))
	if err != nil {
		return err
	}
	keyID, err := hkp.ParseKeyID(ctx.String("recv-keys"))
	if err != nil {
		return err
	}
	client := hkp.NewClient(keyserver, nil)
	keys, err := client.GetKeysByID(context.Background(), keyID)
	if err != nil {
		return err
	}
	return env.SaveKey(keys)
}
