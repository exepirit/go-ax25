package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"log/slog"
	"os"
)

var app = &cli.App{
	Name: "kissmonitor",
	Commands: []*cli.Command{
		monitorCmd,
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "device",
			Required: true,
			Aliases:  []string{"d"},
		},
		&cli.IntFlag{
			Name:    "baudrate",
			Value:   115200,
			Aliases: []string{"b"},
		},
	},
	Before: func(ctx *cli.Context) error {
		slog.SetDefault(slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		))
		return nil
	},
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
