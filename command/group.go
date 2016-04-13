package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/molsbee/clc-cli/api"
)

// GroupCommand Commands related to Group Functions
func GroupCommand() cli.Command {
	return cli.Command{
		Name:  "group",
		Usage: "Provide API Functionality for Group APIs",
		Subcommands: []cli.Command{
			details(),
			getGroupServers(),
		},
	}
}

func details() cli.Command {
	return cli.Command{
		Name:  "details",
		Usage: "Retuner details of Data Center Groups",
		Action: func(ctx *cli.Context) {
			groupID := ctx.Args().Get(0)
			if groupID == "" {
				cli.ShowCommandHelp(ctx, "details")
				return
			}

			group := api.Group{
				ID: groupID,
			}

			details := group.Get()
			fmt.Printf("ID: %s\n", details.ID)
			fmt.Printf("Name: %s\n", details.Name)
			fmt.Printf("Server Count: %d\n", details.ServersCount)
			for _, group := range details.Groups {
				fmt.Println("Group")
				fmt.Printf("     ID:           %s\n", group.ID)
				fmt.Printf("     Name:         %s\n", group.Name)
				fmt.Printf("     Server Count: %d\n", group.ServersCount)
				fmt.Printf("     Servers\n")
				for _, link := range group.Links {
					if link.Rel == "server" {
						fmt.Printf("            ID: %s\n", link.ID)
					}
				}

			}

		},
	}
}

func getGroupServers() cli.Command {
	return cli.Command{
		Name:   "servers",
		Usage:  "Return Servers within Data Center Group",
		Action: func(ctx *cli.Context) {},
	}
}
