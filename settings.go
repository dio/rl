package settings

import "github.com/kelseyhightower/envconfig"

type Settings struct {
	// Host.
	Host string `envconfig:"HOST" default:"0.0.0.0"`
}

func NewSettings() Settings {
	var s Settings
	if err := envconfig.Process("", &s); err != nil {
		panic(err)
	}
	return s
}
