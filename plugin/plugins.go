package plugin

import "time"

// Plugin provides interfaces to handle plugin.
type Plugin interface {
	Start() error
	Stop() error
}

// NewPlugin creates plugin instance
func NewPlugin() Plugin {
	// TODO: Load interval from config.
	return &heimdallPlugin{CheckInterval: time.Second * 5}
}
