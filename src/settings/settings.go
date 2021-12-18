package settings

import (
	"fmt"
	"github.com/spf13/viper"
)

type HostSettings struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
}

type AppSettings struct {
	HostSettings HostSettings `mapstructure:"host"`
	Environment  string       `mapstructure:"env"`
}

func LoadConfig() (settings *AppSettings, err error) {
	// Configure the prefix for the env to avoid collisions
	viper.SetEnvPrefix("task_solver")

	// Set-up default values in case of not found
	viper.SetDefault("env", "dev")

	// Read values using GET from env as well
	viper.AutomaticEnv()

	// Get the current environment (defaults to dev)
	env := viper.GetString("env")

	// Load the respective configuration file
	viper.SetConfigName(fmt.Sprintf("appsettings.%s", env))
	viper.AddConfigPath("./env")
	viper.SetConfigType("json")

	// Check for errors when reading the file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("could not read the appsettings file")
	}

	// Allocate memory for settings
	settings = new(AppSettings)

	// Check for errors during decoding
	if err := viper.Unmarshal(settings); err != nil {
		return nil, fmt.Errorf("could not parse appsettings into struct")
	}

	// Return the settings
	return
}

func (h *HostSettings) GetHost() string {
	return fmt.Sprintf("%s:%s", h.Address, h.Port)
}
