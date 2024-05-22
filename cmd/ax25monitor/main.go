package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var app = &cli.App{
	Name: "ax25monitor",
	Commands: []*cli.Command{
		sendCmd,
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "callsign",
			Required: true,
			Aliases:  []string{"c"},
		},
	},
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
