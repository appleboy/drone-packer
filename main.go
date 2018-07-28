package main

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
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
		cli.StringSliceFlag{
			Name:   "actions",
			Usage:  "a list of actions to have packer perform",
			EnvVar: "PLUGIN_ACTIONS",
		},
		cli.StringFlag{
			Name:   "vars",
			Usage:  "a map of variables to pass to the Packer `build` commands. Each value is passed as a `<key>=<value>` option",
			EnvVar: "PLUGIN_VARS",
		},
		cli.StringSliceFlag{
			Name:   "var_files",
			Usage:  "a list of var files to use. Each value is passed as -var-file=<value>",
			EnvVar: "PLUGIN_VAR_FILES",
		},
		cli.StringSliceFlag{
			Name:   "except",
			Usage:  "validate or build all builds other than these",
			EnvVar: "PLUGIN_EXCEPT",
		},
		cli.StringSliceFlag{
			Name:   "only",
			Usage:  "validate or build only the specified builds",
			EnvVar: "PLUGIN_ONLY",
		},
		cli.StringFlag{
			Name:   "template",
			Usage:  "Will execute multiple builds in parallel as defined in the template",
			EnvVar: "PLUGIN_TEMPLATE",
		},
	}

	app.Version = Version

	if BuildNum != "" {
		app.Version = app.Version + "+" + BuildNum
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	var version string
	if BuildNum != "" {
		version = Version + "+" + BuildNum
	} else {
		version = Version
	}
	logrus.WithFields(logrus.Fields{
		"revision": version,
	}).Info("Drone Packer Plugin Version")

	var vars map[string]string
	if c.String("vars") != "" {
		if err := json.Unmarshal([]byte(c.String("vars")), &vars); err != nil {
			logrus.Panic(err)
		}
	}

	plugin := Plugin{
		Config: Config{
			Actions:  c.StringSlice("actions"),
			Vars:     vars,
			Template: c.String("template"),
			VarFiles: c.StringSlice("var_files"),
			Except:   c.StringSlice("except"),
			Only:     c.StringSlice("only"),
		},
	}

	return plugin.Exec()
}
