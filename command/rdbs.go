package command

import (
	"github.com/molsbee/clc-cli/api"
	"github.com/urfave/cli"
	"fmt"
)

func RdbsCommand() cli.Command {
	return cli.Command{
		Name:  "rdbs",
		Usage: "Collection of RDBS APIs",
		Subcommands: []cli.Command{
			getSubscription(),
		},
	}
}

func getSubscription() cli.Command {
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
			subscriptionId := ctx.String("subscriptionId")

			rdbsSubscription := api.RdbsSubscription{
				BaseRequest: api.BaseRequest{
					AccountAlias: accountAlias,
				},
				SubscriptionId: subscriptionId,
			}

			if subscriptionId != "" {
				subscription := rdbsSubscription.Get()
				fmt.Printf("%+v", subscription)
			} else {
				subscriptions := rdbsSubscription.GetAll()
				for _, subscription := range subscriptions {
					fmt.Printf("%+v", subscription)
				}
			}

			return nil
		},
	}
}
