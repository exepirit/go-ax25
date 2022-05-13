package main

import (
	"flag"
	"log"
	"os"

	"github.com/exepirit/go-ax25/pkg/kiss"
)

func main() {
	out := flag.String("out", "/dev/stdout", "KISS TNC device")
	flag.Parse()

	outStream, err := os.OpenFile(*out, os.O_WRONLY, 0)
	if err != nil {
		log.Fatalln("Cannot open output:", err)
	}
	defer outStream.Close()
	frameWriter := kiss.NewFrameWriter(outStream)

	frame := kiss.Frame{
		Port:    0,
		Command: kiss.DataFrameCommand,
		Data:    []byte("~ HELLO-WORLD! ~"),
	}

	if err = frameWriter.Write(frame); err != nil {
		log.Fatalln("Cannot write frame:", err)
	}
}
