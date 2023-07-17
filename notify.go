package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

func PostSlackWebhook(c *cli.Context, cfg *Config, text string) error {
	webhook := cfg.SlackWebhook

	// only notify when a webhook was provided
	if len(webhook) == 0 {
		logrus.Warnf("no slack webhook provided, skipping notification")
		return nil
	}

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewSectionBlock(
				slack.NewTextBlockObject(slack.MarkdownType, text, false, false),
				nil, nil),
		},
	}

	// construct the message
	msg := slack.WebhookMessage{
		Blocks: &blocks,
	}

	// post to slack
	err := slack.PostWebhook(webhook, &msg)
	if err != nil {
		return fmt.Errorf("unable to post webhook message: %w", err)
	}

	return nil
}
