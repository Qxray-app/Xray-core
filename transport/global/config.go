package global

import (
	"github.com/qxray-app/xray-core/transport/internet"
)

// Apply applies this Config.
func (c *Config) Apply() error {
	if c == nil {
		return nil
	}
	return internet.ApplyGlobalTransportSettings(c.TransportSettings)
}
