package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var app = &cli.App{
	Name:  "ax25monitor",
	Usage: "AX.25 monitor at low-level",
	Commands: []*cli.Command{
		sendCmd,
	},
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
