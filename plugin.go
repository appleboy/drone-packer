package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

type (
	// Config for the plugin.
	Config struct {
		Actions    []string
		Vars       map[string]string
		Template   string
		VarFiles   []string
		SyntaxOnly bool
		Except     []string
		Only       []string
		Color      bool
		Debug      bool
		Parallel   bool
	}

	// Plugin values
	Plugin struct {
		Config Config
	}
)

func pkValidate(config Config) *exec.Cmd {
	args := []string{
		"validate",
	}

	for _, v := range config.VarFiles {
		args = append(args, "-var-file", fmt.Sprintf("%s", v))
	}

	for k, v := range config.Vars {
		args = append(args, "-var")
		args = append(args, fmt.Sprintf("%s=%s", k, v))
	}

	if len(config.Except) > 0 {
		args = append(args, "-except="+strings.Join(config.Except, ","))
	}

	if len(config.Only) > 0 {
		args = append(args, "-only="+strings.Join(config.Only, ","))
	}

	if config.SyntaxOnly {
		args = append(args, "-syntax-only")
	}

	args = append(args, config.Template)
	return exec.Command(
		"packer",
		args...,
	)
}

func pkBuild(config Config) *exec.Cmd {
	args := []string{
		"build",
	}

	for _, v := range config.VarFiles {
		args = append(args, "-var-file", fmt.Sprintf("%s", v))
	}

	for k, v := range config.Vars {
		args = append(args, "-var")
		args = append(args, fmt.Sprintf("%s=%s", k, v))
	}

	if len(config.Except) > 0 {
		args = append(args, "-except="+strings.Join(config.Except, ","))
	}

	if len(config.Only) > 0 {
		args = append(args, "-only="+strings.Join(config.Only, ","))
	}

	if config.Parallel {
		args = append(args, "-parallel=true")
	}

	if config.Color {
		args = append(args, "-color=true")
	}

	if config.Debug {
		args = append(args, "-debug")
	}

	args = append(args, config.Template)
	return exec.Command(
		"packer",
		args...,
	)
}

// Exec executes the plugin.
func (p *Plugin) Exec() error {
	var commands []*exec.Cmd
	if p.Config.Template == "" {
		return errors.New("you must provide a template file")
	}

	if len(p.Config.Actions) == 0 {
		return errors.New("you must provide packer action")
	}

	commands = append(commands, exec.Command("packer", "version"))

	// Add commands listed from Actions
	for _, action := range p.Config.Actions {
		switch action {
		case "validate":
			commands = append(commands, pkValidate(p.Config))
		case "build":
			commands = append(commands, pkBuild(p.Config))
		default:
			return fmt.Errorf("valid actions are: validate, build  You provided %s", action)
		}
	}

	for _, c := range commands {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		err := c.Run()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
			}).Fatal("Failed to execute a command")
		}
		logrus.Debug("Command completed successfully")
	}

	return nil
}
