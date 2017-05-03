package command

import (
	"fmt"

	"github.com/howeyc/gopass"
	"github.com/molsbee/clc-cli/service/clc"
	"github.com/urfave/cli"
)

// Login Commands related to Authenticating with CLC
func Login(api *clc.API) cli.Command {
	return cli.Command{
		Name:        "login",
		Usage:       "Authenticate with CLC V2 API",
		ArgsUsage:   "username",
		Description: "password will be request in secure prompt",
		Action: func(ctx *cli.Context) error {
			username := ctx.Args().Get(0)
			if username == "" {
				return cli.ShowCommandHelp(ctx, "login")
			}

			fmt.Print("Password: ")
			password, _ := gopass.GetPasswd()

			err := api.Authenticate(username, string(password))
			if err != nil {
				fmt.Print(err)
			}

			fmt.Printf("Successfully authentication user %s", username)
			return nil
		},
	}
}
