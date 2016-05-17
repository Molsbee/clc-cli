package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/molsbee/clc-cli/api"
	"github.com/molsbee/clc-cli/json"
)

// ServerCommand Commands related to server functions
func ServerCommand() cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Provide API Functionality for Server APIs",
		Subcommands: []cli.Command{
			detailsCommand(),
			credentialsCommand(),
		},
	}
}

func detailsCommand() cli.Command {
	return cli.Command{
		Name:  "details",
		Usage: "Return Server Details",
		Action: func(ctx *cli.Context) error {
			serverAlias := ctx.Args().Get(0)
			if serverAlias == "" {
				return cli.ShowCommandHelp(ctx, "details")
			}

			server := api.Server{
				Name: serverAlias,
			}

			details := server.Get()
			fmt.Println(json.PrettyPrint(details))
			return nil
		},
	}
}

func credentialsCommand() cli.Command {
	return cli.Command{
		Name:  "credentials",
		Usage: "Return Server Credentials",
		Action: func(ctx *cli.Context) error {
			serverAlias := ctx.Args().Get(0)
			if serverAlias == "" {
				return cli.ShowCommandHelp(ctx, "credentials")
			}

			server := api.Server{
				Name: serverAlias,
			}

			credentials := server.GetCredentials()
			fmt.Println(json.PrettyPrint(credentials))
			return nil
		},
	}
}
