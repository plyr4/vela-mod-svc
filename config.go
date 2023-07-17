package main

import (
	"github.com/urfave/cli/v2"
)

// runtime variables, configured using the environment
type Config struct {
	// variables provided during runtime
	VelaAddr  string
	VelaEnv   string
	VelaToken string

	ServerPort   string
	SlackWebhook string
}

// NewConfig configures and returns a new runtime config
func NewConfig(c *cli.Context) (*Config, error) {
	cfg := Config{
		VelaAddr:  c.String("vela.addr"),
		VelaEnv:   c.String("vela.env"),
		VelaToken: c.String("vela.token"),

		ServerPort:   c.String("server.port"),
		SlackWebhook: c.String("slack.webhook"),
	}

	return &cfg, cfg.validate()
}

// validate validates the user-provided values of the runtime config
func (cfg *Config) validate() error {
	return nil
}
