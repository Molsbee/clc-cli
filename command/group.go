package command

import (
	"fmt"

	"github.com/molsbee/clc-cli/service/clc"
	"github.com/urfave/cli"
)

const GROUP = "group"

// GroupCommand Commands related to Group Functions
func GroupCommand(api *clc.API) cli.Command {
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "accountAlias",
			Usage: "[clc customer account alias]",
		},
	}

	return cli.Command{
		Name:  "groups",
		Usage: "Provide API Functionality for Group APIs",
		Subcommands: []cli.Command{
			details(api, flags),
		},
	}
}

func details(api *clc.API, flags []cli.Flag) cli.Command {
	comName := "get"
	return cli.Command{
		Name:      comName,
		Usage:     "Returns details of Data Center Groups",
		Flags: flags,
		ArgsUsage: "group id or datacenter name",
		Action: func(ctx *cli.Context) error {
			groupID := ctx.Args().Get(0)
			if groupID == "" {
				return cli.ShowCommandHelp(ctx, comName)
			}

			if len(groupID) == 3 {
				dataCenterDetails := api.GetDataCenter(clc.DataCenterRequest{
					Name: groupID,
					AccountAlias: ctx.String("accountAlias"),
				})
				for _, element := range dataCenterDetails.Links {
					if element.Rel == GROUP {
						groupID = element.ID
						break
					}
				}
			}

			fmt.Println(api.GetGroup(clc.GroupRequest{
				ID: groupID,
				AccountAlias: ctx.String("accountAlias"),
			}))
			return nil
		},
	}
}
