package config

import (
	"github.com/spf13/viper"
	"time"
)

type ServerConfig struct {
	Port         string `mapstructure:"port"`
	ReadTimeout  int32  `mapstructure:"read_timeout"`
	WriteTimeout int32  `mapstructure:"write_timeout"`
	MaxHeader    int32  `mapstructure:"max_header"`
}

type Config struct {
	SrvConfig  ServerConfig  `mapstructure:"server"`
	CtxTimeout time.Duration `mapstructure:"timeout"`
}

func Read(configFilePath string) (*Config, error) {
	viper.AutomaticEnv()
	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
