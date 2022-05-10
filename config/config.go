package config

import (
	"time"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port         string `mapstructure:"port"`
	ReadTimeout  int32  `mapstructure:"read_timeout"`
	WriteTimeout int32  `mapstructure:"write_timeout"`
	MaxHeader    int32  `mapstructure:"max_header"`
}

type DbConfig struct {
	Uri string `mapstructure:"uri"`
}

type Config struct {
	SrvConfig      ServerConfig  `mapstructure:"server"`
	DatabaseConfig DbConfig      `mapstructure:"db"`
	CtxTimeout     time.Duration `mapstructure:"timeout"`
	SecretJWT      string        `mapstructure:"secret"`
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
