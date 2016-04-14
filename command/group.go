package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/molsbee/clc-cli/api"
)

const server = "server"

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
			fmt.Printf("ID:\t\t%s\n", details.ID)
			fmt.Printf("Name:\t\t%s\n", details.Name)
			fmt.Printf("Server Count:\t%d\n", details.ServersCount)
			for _, group := range details.Groups {
				fmt.Println("Group")
				fmt.Printf("\tID:\t\t%s\n", group.ID)
				fmt.Printf("\tName:\t\t%s\n", group.Name)
				fmt.Printf("\tServer Count:\t%d\n", group.ServersCount)
				if group.ServersCount != 0 {
					fmt.Printf("\tServers\n")
				}
				for _, link := range group.Links {
					if link.Rel == server {
						fmt.Printf("\t\tID: %s\n", link.ID)
					}
				}

			}

		},
	}
}

func getGroupServers() cli.Command {
	return cli.Command{
		Name:  "servers",
		Usage: "Return Servers within Data Center Group",
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

			fmt.Println("Servers")
			for _, link := range details.Links {
				if link.Rel == server {
					fmt.Printf("\tID: %s\n", link.ID)
				}
			}
			for _, group := range details.Groups {
				for _, link := range group.Links {
					if link.Rel == server {
						fmt.Printf("\tID: %s\n", link.ID)
					}
				}
			}
		},
	}
}
