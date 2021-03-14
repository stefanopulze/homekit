package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Current cfg

func Init() (*cfg, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	viper.SetEnvPrefix("homekit")
	viper.AutomaticEnv()

	viper.AddConfigPath("/etc/homekit/")
	viper.AddConfigPath("$HOME/.homekit")
	viper.AddConfigPath(".")

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&Current); err != nil {
		return nil, err
	}

	logrus.Tracef("Config: %+v", Current)

	return &Current, nil
}

func setDefaults() {
	viper.SetDefault("homekit.name", "dev")
	viper.SetDefault("homekit.pin", "00102003")
	viper.SetDefault("homekit.storagePath", "./db")
	viper.SetDefault("server.port", 4000)
}
