module github.com/exepirit/go-ax25/cmd/ax25monitor

go 1.22.3

require (
	github.com/exepirit/go-ax25 v0.0.0
	github.com/urfave/cli/v2 v2.27.2
)

replace github.com/exepirit/go-ax25 v0.0.0 => ../..

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240312152122-5f08fbb34913 // indirect
)
