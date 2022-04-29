package config

type ServerConfig struct {
	Port         string `mapstructure:"port"`
	ReadTimeout  int32  `mapstructure:"read_timeout"`
	WriteTimeout int32  `mapstructure:"write_timeout"`
	MaxHeader    int32  `mapstructure:"max_header"`
}
