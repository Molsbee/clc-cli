package command

import (
	"fmt"

	"github.com/molsbee/clc-cli/service/clc"
	"github.com/urfave/cli"
)

// ServerCommand Commands related to server functions
func ServerCommand(api *clc.API) cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Provide API Functionality for Server APIs",
		Subcommands: []cli.Command{
			detailsCommand(api),
			credentialsCommand(api),
		},
	}
}

func detailsCommand(api *clc.API) cli.Command {
	return cli.Command{
		Name:  "details",
		Usage: "Return Server Details",
		Action: func(ctx *cli.Context) error {
			serverAlias := ctx.Args().Get(0)
			if serverAlias == "" {
				return cli.ShowCommandHelp(ctx, "details")
			}

			fmt.Println(api.GetServer(clc.ServerRequest{Name: serverAlias}))
			return nil
		},
	}
}

func credentialsCommand(api *clc.API) cli.Command {
	return cli.Command{
		Name:  "credentials",
		Usage: "Return Server Credentials",
		Action: func(ctx *cli.Context) error {
			serverAlias := ctx.Args().Get(0)
			if serverAlias == "" {
				return cli.ShowCommandHelp(ctx, "credentials")
			}

			fmt.Println(api.GetServerCredentials(clc.ServerRequest{Name: serverAlias}))
			return nil
		},
	}
}
