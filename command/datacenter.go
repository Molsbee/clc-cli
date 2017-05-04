package command

import (
	"fmt"

	"github.com/molsbee/clc-cli/service/clc"
	"github.com/urfave/cli"
)

// TODO: Update to accept AccountAlias as optional parameter
// DataCenterCommand Commands related to Data Center functions
func DataCenterCommand(api *clc.API) cli.Command {
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "accountAlias",
			Usage: "[clc customer account alias]",
		},
	}

	return cli.Command{
		Name:  "datacenter",
		Usage: "command line wrapper for retrieving datacenter data",
		Subcommands: []cli.Command{
			get(api, flags),
			list(api, flags),
		},
	}
}

func get(api *clc.API, flags []cli.Flag) cli.Command {
	comName := "get"
	return cli.Command{
		Name:  comName,
		Usage: "load data center details",
		Flags: flags,
		Action: func(ctx *cli.Context) error {
			name := ctx.Args().Get(0)
			if name == "" {
				return cli.ShowCommandHelp(ctx, comName)
			}

			fmt.Print(api.GetDataCenter(clc.DataCenterRequest{
				Name: name,
				AccountAlias: ctx.String("accountAlias"),
			}))
			return nil
		},
	}
}

func list(api *clc.API, flags []cli.Flag) cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "load all datacenters",
		Flags: flags,
		Action: func(ctx *cli.Context) error {
			fmt.Print(api.GetDataCenters(clc.DataCenterRequest{}))
			return nil
		},
	}
}
