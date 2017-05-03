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
		Usage: "Returns subscriptions associated to account alias",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "accountAlias",
				Usage: "customer account alias",
			},
			cli.StringFlag{
				Name:  "subscriptionId",
				Usage: "Subscription ID that needs details",
			},
		},
		Action: func(ctx *cli.Context) error {
			accountAlias := ctx.String("accountAlias")
			subscriptionIdString := ctx.String("subscriptionId")

			subscriptionId, _ := strconv.Atoi(subscriptionIdString)
			request := rdbs.SubscriptionRequest{
				AccountAlias:   accountAlias,
				SubscriptionID: subscriptionId,
			}

			if subscriptionIdString != "" {
				fmt.Println(api.GetSubscription(request))
			} else {
				fmt.Println(api.GetSubscriptions(request))
			}

			return nil
		},
	}
}
