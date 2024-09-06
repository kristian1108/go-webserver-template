package config

import (
	"strings"
)
import "github.com/spf13/viper"

type Test struct {
	Test string
}

type TestConfig struct {
	Test Test
}

func Load(path string) (TestConfig, error) {
	var c TestConfig
	viper.AddConfigPath(path)
	viper.SetConfigName("Config")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		return c, err
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("test")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.Unmarshal(&c)
	if err != nil {
		return c, err
	}

	return c, nil
}
