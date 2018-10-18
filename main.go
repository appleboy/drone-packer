package main

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"
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
		cli.BoolFlag{
			Name:   "syntax_only",
			Usage:  "Only check syntax. Do not verify config of the template",
			EnvVar: "PLUGIN_SYNTAX_ONLY",
		},
		cli.BoolFlag{
			Name:   "color",
			Usage:  "Disable color output (on by default)",
			EnvVar: "PLUGIN_COLOR",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Debug mode enabled for builds",
			EnvVar: "PLUGIN_DEBUG",
		},
		cli.BoolFlag{
			Name:   "parallel",
			Usage:  "Disable parallelization (on by default)",
			EnvVar: "PLUGIN_PARALLEL",
		},
		cli.BoolFlag{
			Name:   "readable",
			Usage:  "Machine-readable output",
			EnvVar: "PLUGIN_READABLE",
		},
		cli.BoolFlag{
			Name:   "force",
			Usage:  "Force a build to continue if artifacts exist, deletes existing artifacts",
			EnvVar: "PLUGIN_FORCE",
		},
	}

	app.Version = Version

	if BuildNum != "" {
		app.Version = app.Version + "+" + BuildNum
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("app can't rum")
	}
}

func run(c *cli.Context) error {
	var version string
	if BuildNum != "" {
		version = Version + "+" + BuildNum
	} else {
		version = Version
	}
	log.Info().Str("revision", version).Msg("Drone Packer Plugin Version")

	var vars map[string]string
	if c.String("vars") != "" {
		if err := json.Unmarshal([]byte(c.String("vars")), &vars); err != nil {
			log.Fatal().Err(err).Msg("json unmarshal")
		}
	}

	plugin := Plugin{
		Config: Config{
			Actions:    c.StringSlice("actions"),
			Vars:       vars,
			Template:   c.String("template"),
			VarFiles:   c.StringSlice("var_files"),
			Except:     c.StringSlice("except"),
			Only:       c.StringSlice("only"),
			SyntaxOnly: c.Bool("syntax_only"),
			Color:      c.Bool("color"),
			Debug:      c.Bool("debug"),
			Parallel:   c.Bool("parallel"),
			Readable:   c.Bool("readable"),
			Force:      c.Bool("force"),
		},
	}

	return plugin.Exec()
}
