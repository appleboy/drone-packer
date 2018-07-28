package main

type (
	// Config for the plugin.
	Config struct {
	}

	// Plugin values
	Plugin struct {
		Config Config
	}
)

// Exec executes the plugin.
func (p *Plugin) Exec() error {
	return nil
}
