package main

import (
	"fmt"
	"github.com/exepirit/go-ax25/ax25"
	"github.com/exepirit/go-ax25/kiss"
	"github.com/urfave/cli/v2"
	"go.bug.st/serial"
)

var sendCmd = &cli.Command{
	Name:  "send",
	Usage: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "device",
			Usage:    "",
			Required: true,
			Aliases:  []string{"d"},
		},
		&cli.IntFlag{
			Name:    "baudrate",
			Usage:   "",
			Value:   115200,
			Aliases: []string{"b"},
		},
	},
	Action: func(ctx *cli.Context) error {
		mode := &serial.Mode{
			BaudRate: ctx.Int("baudrate"),
		}
		port, err := serial.Open(ctx.String("device"), mode)
		if err != nil {
			return fmt.Errorf("failed to open serial port: %w", err)
		}
		defer port.Close()

		kissWriter := kiss.NewEncoder(port, 0)
		ax25Writer := ax25.NewPacketWriter(kissWriter, 256)

		packet := ax25.Packet{
			Address: ax25.PacketAddress{
				Destination: ax25.MustParseAddress("NOCALL-1"),
				Source:      ax25.MustParseAddress("NOCALL-2"),
			},
			Control: ax25.ControlData{
				Type:    ax25.PacketTypeUnnumbered,
				IsFinal: false,
			},
			PID:  ax25.ProtocolNoLayer3,
			Info: []byte("test"),
		}
		err = ax25Writer.Write(&packet)
		return err
	},
}
