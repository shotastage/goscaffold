package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Go Scaffold"
	app.Usage = "Multi scaffold"
	app.Version = "0.0.1"

	app.Run(os.Args)
}
