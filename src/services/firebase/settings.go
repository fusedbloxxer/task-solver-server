package firebase

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"os"
)

type Credentials struct {
	Key    string `mapstructure:"key"`
	Secret string `mapstructure:"secret"`
}

type Settings struct {
	Credentials Credentials `mapstructure:"credentials"`
}

func readSettings(key string, settings map[string]interface{}) (*Settings, error) {
	var exists bool
	var config interface{}

	if config, exists = settings[SectionKey]; !exists {
		return nil, fmt.Errorf("could not read settings for %s", key)
	}

	var firebaseSettings Settings
	if err := mapstructure.Decode(config, &firebaseSettings); err != nil {
		return nil, fmt.Errorf("invalid conversion to settings: %w", err)
	}

	return &firebaseSettings, nil
}

func (c *Credentials) load() (err error) {
	if err = os.Setenv(c.Key, c.Secret); err != nil {
		return fmt.Errorf("could not set firebase credentials: %w", err)
	}

	return
}
