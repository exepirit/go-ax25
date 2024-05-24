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
	Usage: "Send data in UI packet",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "device",
			Usage:    "serial device name or path",
			Required: true,
			Aliases:  []string{"d"},
		},
		&cli.IntFlag{
			Name:    "baudrate",
			Usage:   "serial device baudrate",
			Value:   115200,
			Aliases: []string{"b"},
		},
		&cli.StringFlag{
			Name:     "source",
			Usage:    "source callsign in CALLSIGN-SSID format",
			Aliases:  []string{"src"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "destination",
			Usage:    "destination callsign in CALLSIGN-SSID format",
			Aliases:  []string{"dst", "dest"},
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		srcAddress, err := ax25.ParseAddress(ctx.String("source"))
		if err != nil {
			return fmt.Errorf("invalid source callsign: %w", err)
		}
		dstAddress, err := ax25.ParseAddress(ctx.String("destination"))
		if err != nil {
			return fmt.Errorf("invalid destination callsign: %w", err)
		}

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
				Destination: dstAddress,
				Source:      srcAddress,
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
