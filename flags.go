package main

import (
	"github.com/urfave/cli/v2"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		EnvVars: []string{"LOG_LEVEL"},
		Name:    "log.level",
		Usage:   "application logging level",
		Value:   "info",
	},
	&cli.StringFlag{
		EnvVars: []string{"SERVER_PORT"},
		Name:    "server.port",
		Usage:   "modification service port number",
		Value:   "8089",
	},
	&cli.StringFlag{
		EnvVars: []string{"VELA_ADDRESS", "VELA_ADDR"},
		Name:    "vela.addr",
		Usage:   "Vela server address as a fully qualified url (<scheme>://<host>)",
		Value:   "https://vela-server.com",
	},
	&cli.StringFlag{
		EnvVars: []string{"VELA_ENVIRONMENT", "VELA_ENV"},
		Name:    "vela.env",
		Usage:   "Vela environment (local|dev|prod)",
		Value:   "local",
	},
	&cli.StringFlag{
		EnvVars: []string{"VELA_TOKEN"},
		Name:    "vela.token",
		Usage:   "Vela platform admin token used for registering workers",
	},
	&cli.StringFlag{
		EnvVars: []string{"SLACK_WEBHOOK"},
		Name:    "slack.webhook",
		Usage:   "The Slack webhook destination for posting notifications. If empty, notifications will not be sent",
		Value:   "",
	},
}
