package command

import (
	"fmt"

	"github.com/molsbee/clc-cli/api"
	"github.com/urfave/cli"
)

// DataCenterCommand Commands related to Data Center functions
func DataCenterCommand() cli.Command {
	return cli.Command{
		Name:  "datacenter",
		Usage: "Provide API Functionality for Data Center APIs",
		Subcommands: []cli.Command{
			get(),
			list(),
		},
	}
}

func get() cli.Command {
	return cli.Command{
		Name:  "details",
		Usage: "Return Data Center Details",
		Action: func(ctx *cli.Context) error {
			name := ctx.Args().Get(0)
			if name == "" {
				return cli.ShowCommandHelp(ctx, "details")
			}

			fmt.Println(api.DataCenter{Name: name}.Get())
			return nil
		},
	}
}

func list() cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "Return List of Data Centers with Details",
		Action: func(ctx *cli.Context) error {
			fmt.Println(api.DataCenter{}.List())
			return nil
		},
	}
}
