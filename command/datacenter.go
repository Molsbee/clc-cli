package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/molsbee/clc-cli/api"
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
		Action: func(ctx *cli.Context) {
			name := ctx.Args().Get(0)
			if name == "" {
				cli.ShowCommandHelp(ctx, "details")
				return
			}

			fmt.Println(api.DataCenter{Name: name}.Get())
		},
	}
}

func list() cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "Return List of Data Centers with Details",
		Action: func(ctx *cli.Context) {
			fmt.Println(api.DataCenter{}.List())
		},
	}
}
