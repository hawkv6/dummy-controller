package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	viperInstance = viper.NewWithOptions(viper.KeyDelimiter("\\"))
	Params        Config
)

type ServiceConfig struct {
	Intent string
	Sid    []string
}

type Config struct {
	Address  string
	Port     int
	Services map[string][]ServiceConfig
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
