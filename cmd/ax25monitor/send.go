package main

import (
	"errors"
	"github.com/exepirit/go-ax25/pkg/ax25"
	"github.com/urfave/cli/v2"
)

var sendCmd = &cli.Command{
	Name:      "send",
	Args:      true,
	ArgsUsage: "data",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "remote",
			Aliases:  []string{"r"},
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		localAddr, err := ax25.ParseAddress(ctx.String("callsign"))
		if err != nil {
			return errors.New("invalid callsign value")
		}
		remoteAddr, err := ax25.ParseAddress(ctx.String("remote"))
		if err != nil {
			return errors.New("invalid remote callsign value")
		}

		conn, err := ax25.DialUnnumbered(&localAddr, &remoteAddr)
		if err != nil {
			return err
		}
		defer conn.Close()

		data := []byte(ctx.Args().First())
		_, err = conn.Write(data)
		return err
	},
}
