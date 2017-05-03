package command

import (
	"fmt"

	"github.com/molsbee/clc-cli/service/clc"
	"github.com/urfave/cli"
)

const server = "server"

// GroupCommand Commands related to Group Functions
func GroupCommand(api *clc.API) cli.Command {
	return cli.Command{
		Name:  "group",
		Usage: "Provide API Functionality for Group APIs",
		Subcommands: []cli.Command{
			details(api),
			//getGroupServers(clc),
		},
	}
}

func details(api *clc.API) cli.Command {
	return cli.Command{
		Name:      "details",
		Usage:     "Returns details of Data Center Groups",
		ArgsUsage: "group id or datacenter name",
		Action: func(ctx *cli.Context) error {
			groupID := ctx.Args().Get(0)
			if groupID == "" {
				return cli.ShowCommandHelp(ctx, "details")
			}

			if len(groupID) == 3 {
				dataCenterDetails := api.GetDataCenter(clc.DataCenterRequest{Name: groupID})
				for _, element := range dataCenterDetails.Links {
					if element.Rel == "group" {
						groupID = element.ID
						break
					}
				}
			}

			fmt.Println(api.GetGroup(clc.GroupRequest{ID: groupID}))
			return nil
		},
	}
}

//func getGroupServers(clc *clc.API) cli.Command {
//	return cli.Command{
//		Name:  "servers",
//		Usage: "Return Servers within Data Center Group",
//		Flags: []cli.Flag{
//			cli.StringFlag{
//				Name:  "data-center",
//				Usage: "Data Center Name ex: UC1",
//			},
//			cli.StringFlag{
//				Name:  "group-id",
//				Usage: "Group ID that needs details",
//			},
//		},
//		Action: func(ctx *cli.Context) error {
//			groupID := ctx.String("group-id")
//
//			dataCenter := ctx.String("data-center")
//			if dataCenter != "" {
//				dataCenterDetails := api.DataCenter{Name: dataCenter}
//				for _, element := range dataCenterDetails.Get().Links {
//					if element.Rel == "group" {
//						groupID = element.ID
//						break
//					}
//				}
//			}
//
//			if groupID == "" {
//				return cli.ShowCommandHelp(ctx, "servers")
//			}
//
//			group := api.Group{
//				ID: groupID,
//			}
//
//			fmt.Println("Servers")
//			fmt.Print(group.Get().GetServers())
//			return nil
//		},
//	}
//}
