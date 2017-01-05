package command

import (
	"fmt"

	"github.com/howeyc/gopass"
	"github.com/molsbee/clc-cli/authentication"
	"github.com/urfave/cli"
)

// Login Commands related to Authenticating with CLC
func Login() cli.Command {
	return cli.Command{
		Name:  "login",
		Usage: "Authenticate with CLC V2 API",
		Action: func(ctx *cli.Context) error {
			username := ctx.Args().Get(0)
			if username == "" {
				return cli.ShowCommandHelp(ctx, "login")
			}

			fmt.Printf("Password: ")
			password, _ := gopass.GetPasswd()

			authentication.Authenticate(username, string(password))
			return nil
		},
	}
}
