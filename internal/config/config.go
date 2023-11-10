package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	viperInstance = viper.NewWithOptions(viper.KeyDelimiter("\\"))
	Params        Config
)

type Intent struct {
	Intent string   `mapstructure:"intent"`
	Sid    []string `mapstructure:"sid"`
}

type ServiceConfig struct {
	Ipv6Addresses []string `mapstructure:"ipv6_addresses"`
	Intents       []Intent `mapstructure:"intents"`
}

type Config struct {
	Address  string                   `mapstructure:"address"`
	Port     int                      `mapstructure:"port"`
	Services map[string]ServiceConfig `mapstructure:"services"`
}

func init() {
	viperInstance.SetEnvPrefix("DC")
	viperInstance.AutomaticEnv()
}

func Parse() error {
	if len(viperInstance.ConfigFileUsed()) != 0 {
		if err := viperInstance.ReadInConfig(); err != nil {
			return fmt.Errorf("failed to load config file %s: %v", viperInstance.ConfigFileUsed(), err)
		}
	}

	if err := viperInstance.UnmarshalExact(&Params); err != nil {
		return fmt.Errorf("failed to parse config: %v", err)
	}

	return nil
}

func GetInstance() *viper.Viper {
	return viperInstance
}
