package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/howeyc/gopass"
	"github.com/molsbee/clc-cli/authentication"
)

// Login Commands related to Authenticating with CLC
func Login() cli.Command {
	return cli.Command{
		Name:  "login",
		Usage: "Authenticate with CLC V2 API",
		Action: func(ctx *cli.Context) {
			username := ctx.Args().Get(0)
			if username == "" {
				cli.ShowCommandHelp(ctx, "login")
				return
			}

			fmt.Printf("Password: ")
			password, _ := gopass.GetPasswd()

			authentication.Authenticate(username, string(password))
		},
	}
}
