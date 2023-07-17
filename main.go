package main

import (
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	// set up the app before running
	app.Before = load

	// define how to run the app
	app.Action = run

	// flags
	app.Flags = flags

	app.Name = "vela-mod-svc"
	app.HelpName = "Vela Modification Service"
	app.Usage = "microservice for analyzing/modifying vela pipelines mid-flight"
	app.Copyright = "Copyright (c) 2023 Target Brands, Inc. All rights reserved."
	app.Authors = []*cli.Author{
		{
			Name:  "Vela Admins",
			Email: "vela@target.com",
		},
	}

	app.Compiled = time.Now()
	app.EnableBashCompletion = true
	app.UseShortOptionHandling = true
	app.Version = "v0.0.1"

	// initiate app execution
	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
