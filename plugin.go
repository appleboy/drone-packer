package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type (
	// Config for the plugin.
	Config struct {
		Actions    []string
		Vars       map[string]string
		Template   string
		Context    string
		VarFiles   []string
		SyntaxOnly bool
		Except     []string
		Only       []string
		Color      bool
		Debug      bool
		Parallel   bool
		Readable   bool
		Force      bool
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
		args = append(args, fmt.Sprintf("-var-file=%s", v))
	}

	for k, v := range config.Vars {
		args = append(args, "-var", fmt.Sprintf("%s=%s", k, v))
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

	cmd := exec.Command(
		"packer",
		args...,
	)

	cmd.Dir = config.Context

	return cmd
}

func pkBuild(config Config) *exec.Cmd {
	args := []string{
		"build",
	}

	for _, v := range config.VarFiles {
		args = append(args, fmt.Sprintf("-var-file=%s", v))
	}

	for k, v := range config.Vars {
		args = append(args, "-var", fmt.Sprintf("%s=%s", k, v))
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

	if config.Readable {
		args = append(args, "-machine-readable")
	}

	if config.Force {
		args = append(args, "-force")
	}

	args = append(args, config.Template)

	cmd := exec.Command(
		"packer init;",
		"packer",
		args...,
	)

	cmd.Dir = config.Context

	return cmd
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

	for _, cmd := range commands {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Env = os.Environ()

		trace(cmd)

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}
