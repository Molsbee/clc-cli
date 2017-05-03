package main

import (
	"fmt"
	"os"
	"time"

	"encoding/json"
	"github.com/molsbee/clc-cli/command"
	"github.com/molsbee/clc-cli/config"
	"github.com/molsbee/clc-cli/model"
	"github.com/molsbee/clc-cli/service/clc"
	"github.com/molsbee/clc-cli/service/rdbs"
	"github.com/urfave/cli"
	"io/ioutil"
)

func main() {
	auth := parseConfigFileForAuthDetails()
	tokenDetails, tokenError := auth.DecodeBearerToken()

	clcAPI := clc.NewAPI(auth)
	rdbsAPI := rdbs.NewAPI(auth)

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
		command.Login(clcAPI),
		command.ServerCommand(clcAPI),
		command.DataCenterCommand(clcAPI),
		command.GroupCommand(clcAPI),
		command.SSHCommand(clcAPI),
		command.RdbsCommand(rdbsAPI),
	}

	app.Run(os.Args)
	fmt.Println()
}

func parseConfigFileForAuthDetails() model.AuthenticationResponse {
	fileData, _ := ioutil.ReadFile(config.FilePath())
	auth := model.AuthenticationResponse{}
	json.Unmarshal(fileData, &auth)
	return auth
}
