package command

import (
	"fmt"

	"github.com/molsbee/clc-cli/service/clc"
	"github.com/urfave/cli"
)

// TODO: Update to accept AccountAlias as optional parameter
// DataCenterCommand Commands related to Data Center functions
func DataCenterCommand(api *clc.API) cli.Command {
	return cli.Command{
		Name:  "datacenter",
		Usage: "Provide API Functionality for Data Center APIs",
		Subcommands: []cli.Command{
			get(api),
			list(api),
		},
	}
}

func get(api *clc.API) cli.Command {
	return cli.Command{
		Name:  "details",
		Usage: "Return Data Center Details",
		Action: func(ctx *cli.Context) error {
			name := ctx.Args().Get(0)
			if name == "" {
				return cli.ShowCommandHelp(ctx, "details")
			}

			fmt.Println(api.GetDataCenter(clc.DataCenterRequest{Name: name}))
			return nil
		},
	}
}

func list(api *clc.API) cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "Return List of Data Centers with Details",
		Action: func(ctx *cli.Context) error {
			fmt.Println(api.GetDataCenters(clc.DataCenterRequest{}))
			return nil
		},
	}
}
