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

			dataCenter := api.DataCenter{
				Name: name,
			}

			details := dataCenter.Get()
			fmt.Printf("ID:\t%s\n", details.ID)
			fmt.Printf("Name:\t%s\n", details.Name)
			for _, element := range details.Links {
				if element.Rel == "group" {
					fmt.Println("Group")
					fmt.Printf("\tName:\t%s\n", element.Name)
					fmt.Printf("\tID:\t%s\n", element.ID)
				}
			}
		},
	}
}

func list() cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "Return List of Data Centers with Details",
		Action: func(ctx *cli.Context) {
			dataCenter := api.DataCenter{}

			dataCenters := dataCenter.List()
			for _, element := range dataCenters {
				fmt.Printf("ID: %s\t\tName: %s\n", element.ID, element.Name)
			}
		},
	}
}
