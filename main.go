package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fanyang89/zerologging/v1"
	"github.com/goccy/go-json"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

var app = &cli.App{
	Name: "slack-status",
	Commands: []*cli.Command{
		{
			Name: "get",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "user-id",
					EnvVars: []string{
						"SLACK_USER_ID",
					},
				},
			},
			Action: func(ctx *cli.Context) error {
				client := slack.New(os.Getenv("SLACK_TOKEN"))
				profile, err := client.GetUserProfile(&slack.GetUserProfileParameters{
					UserID: ctx.String("user-id"),
				})
				if err != nil {
					return errors.Wrap(err, "get user profile failed")
				}

				display, err := json.MarshalIndent(map[string]string{
					"emoji":  profile.StatusEmoji,
					"status": profile.StatusText,
					"expire": strconv.Itoa(profile.StatusExpiration),
				}, "", "  ")
				if err != nil {
					return errors.Wrap(err, "marshal display failed")
				}

				fmt.Println(string(display))
				return nil
			},
		},
		{
			Name: "set",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "status",
					Required: true,
				},
				&cli.StringFlag{
					Name: "emoji",
				},
				&cli.Int64Flag{
					Name:  "expire",
					Value: 0,
				},
			},
			Action: func(ctx *cli.Context) error {
				client := slack.New(os.Getenv("SLACK_TOKEN"))
				err := client.SetUserCustomStatus(
					ctx.String("status"),
					ctx.String("emoji"),
					ctx.Int64("expire"))
				if err != nil {
					return errors.Wrap(err, "set user status failed")
				}
				return nil
			},
		},
	},
}

func main() {
	zerologging.WithConsoleLog(zerolog.InfoLevel)

	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Load env file failed")
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("Unexpected error")
	}
}
