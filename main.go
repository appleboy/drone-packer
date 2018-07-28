package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// Version set at compile-time
var (
	Version  string
	BuildNum string
)

func main() {
	app := cli.NewApp()
	app.Name = "packer plugin"
	app.Usage = "packer plugin"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Usage:  "telegram token",
			EnvVar: "PLUGIN_TOKEN,TELEGRAM_TOKEN",
		},
	}

	app.Version = Version

	if BuildNum != "" {
		app.Version = app.Version + "+" + BuildNum
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{}

	return plugin.Exec()
}
