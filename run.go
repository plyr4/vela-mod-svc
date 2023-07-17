package main

import (
	"github.com/go-vela/sdk-go/vela"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func run(c *cli.Context) error {
	// configure and validate runtime configurations
	cfg, err := NewConfig(c)
	if err != nil {
		logrus.Fatalf("unable to start, invalid configuration: %s", err.Error())
	}

	// create a vela client
	logrus.Info("creating vela client")

	velaClient, err := vela.NewClient(cfg.VelaAddr, "", nil)
	if err != nil {
		logrus.Fatalf("unable to create vela client: %s", err.Error())
	}

	velaClient.Authentication.SetPersonalAccessTokenAuth(cfg.VelaToken)

	// run server indefinitely
	logrus.Info("starting modification service...")

	return server(cfg)
}
