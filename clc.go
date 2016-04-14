package main

import (
	"fmt"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/molsbee/clc-cli/authentication"
	"github.com/molsbee/clc-cli/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "CLC CLI"
	app.Usage = "Helper CLI for accessing CLC Resources"

	app.Before = func(ctx *cli.Context) error {
		command := ctx.Args().First()

		tokenDetails, err := authentication.Auth.DecodeBearerToken()
		if err != nil && command != "login" {
			fmt.Println("top")
			return fmt.Errorf("Please use Login command to authenticate with CLC")
		} else if err == nil {
			experation := time.Unix(tokenDetails.EXP, 0)
			if time.Now().After(experation) {
				return fmt.Errorf("Please use Login command to authenticate with CLC")
			}

			app.UsageText = "\tCurrent User: " + tokenDetails.UniqueName + "\n\tAccount: " + tokenDetails.AccountAlias
		}

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
