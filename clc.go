package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/molsbee/clc-cli/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "CLC CLI"
	app.Usage = "Helper CLI for accessing CLC Resources"
	app.Commands = []cli.Command{command.Login(), command.SSHCommand()}

	app.Run(os.Args)
}
