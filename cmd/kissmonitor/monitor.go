package main

import (
	"fmt"
	"github.com/exepirit/go-ax25/pkg/kiss"
	"github.com/urfave/cli/v2"
	"go.bug.st/serial"
	"log/slog"
)

var monitorCmd = &cli.Command{
	Name:  "monitor",
	Usage: "Monitor KISS frames and its content",
	Action: func(ctx *cli.Context) error {
		mode := &serial.Mode{
			BaudRate: ctx.Int("baudrate"),
		}
		port, err := serial.Open(ctx.String("device"), mode)
		if err != nil {
			return fmt.Errorf("failed to open serial port: %w", err)
		}
		defer port.Close()

		reader := kiss.NewFrameReader(port)
		for {
			frame, err := reader.Read()
			if err != nil {
				slog.Error("read frame error", "error", err)
				continue
			}

			fmt.Printf("New frame (port = %d, cmd = %02x)\n", frame.Port, frame.Command)
			for i := 0; i < len(frame.Data); i += 16 {
				fmt.Printf("%08x  ", i)
				for j := 0; j < 16 && i+j < len(frame.Data); j++ {
					fmt.Printf("%02x ", frame.Data[i+j])
				}
				fmt.Println()
			}
			fmt.Println()
		}
	},
}
