package command

import (
	"fmt"
	"github.com/molsbee/clc-cli/service/rdbs"
	"github.com/urfave/cli"
	"strconv"
)

func RdbsCommand(rdbs *rdbs.API) cli.Command {
	return cli.Command{
		Name:  "rdbs",
		Usage: "Collection of RDBS APIs",
		Subcommands: []cli.Command{
			getSubscription(rdbs),
		},
	}
}

func getSubscription(api *rdbs.API) cli.Command {
	return cli.Command{
		Name:  "subscriptions",
		Description: "returns one or more subscriptions based on input",
		Usage: "optional parameters [--accountAlias {{alias}}] [--subscriptionId {{id}}] [--environment {{ env }}]",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "accountAlias",
				Usage: "customer account alias",
			},
			cli.StringFlag{
				Name:  "subscriptionId",
				Usage: "subscriptoin id that needs details",
			},
			cli.StringFlag{
				Name: "environment",
				Usage: "dev,qa,prod",
			},
		},
		Action: func(ctx *cli.Context) error {
			accountAlias := ctx.String("accountAlias")
			subscriptionIdString := ctx.String("subscriptionId")
			environment := ctx.String("environment")

			subscriptionId, _ := strconv.Atoi(subscriptionIdString)
			request := rdbs.SubscriptionRequest{
				AccountAlias:   accountAlias,
				SubscriptionID: subscriptionId,
				EnvironmentRequest: rdbs.EnvironmentRequest{
					Environment: environment,
				},
			}

			if subscriptionIdString != "" {
				fmt.Print(api.GetSubscription(request))
			} else {
				subscriptions := api.GetSubscriptions(request)
				fmt.Println("Subscriptions")
				for _, sub := range subscriptions {
					fmt.Print(sub.FmtString("\t"))
				}
			}

			return nil
		},
	}
}
