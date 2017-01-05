package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
	"github.com/molsbee/clc-cli/authentication"
	"github.com/molsbee/clc-cli/command"
)

func main() {
	tokenDetails, tokenError := authentication.Auth.DecodeBearerToken()

	app := cli.NewApp()
	app.Name = "CLC CLI"
	app.Usage = "Helper CLI for accessing CLC Resources"

	app.Before = func(ctx *cli.Context) error {
		command := ctx.Args().First()

		if command != "login" {
			if tokenError != nil {
				return fmt.Errorf("Please use Login command to authenticate with CLC")
			}
			experation := time.Unix(tokenDetails.EXP, 0)
			if time.Now().After(experation) {
				return fmt.Errorf("Please use login command to authenticate with CLC bearer token expired %s", experation)
			}
		}

		app.UsageText = "\tCurrent User: " + tokenDetails.UniqueName + "\n\tAccount: " + tokenDetails.AccountAlias
		return nil
	}

	app.Commands = []cli.Command{
		command.Login(),
		command.ServerCommand(),
		command.DataCenterCommand(),
		command.GroupCommand(),
		command.SSHCommand(),
	}

	app.Run(os.Args)
	fmt.Println()
}
