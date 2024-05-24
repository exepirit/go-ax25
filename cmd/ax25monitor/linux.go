package main

import "github.com/urfave/cli/v2"

var linuxCmd = &cli.Command{
	Name:  "linux",
	Usage: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "callsign",
			Required: true,
			Aliases:  []string{"c"},
		},
	},
	Subcommands: []*cli.Command{
		linuxSendCmd,
	},
}
