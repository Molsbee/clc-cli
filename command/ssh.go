package command

import (
	"github.com/molsbee/clc-cli/ssh"
	"github.com/urfave/cli"
)

// SSHCommand Commands related to creating ssh connection to CLC Server
func SSHCommand() cli.Command {
	return cli.Command{
		Name:        "ssh",
		Usage:       "Connect to CLC Server with SSH",
		Description: "Example: clc ssh {{server alias}}",
		Action: func(ctx *cli.Context) error {
			serverAlias := ctx.Args().Get(0)
			if serverAlias == "" {
				return cli.ShowCommandHelp(ctx, "ssh")
			}

			ssh.Connect(serverAlias)
			return nil
		},
	}
}
