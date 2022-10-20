package config

type Redis struct {
	Addr     string `mapstructure:"addr" yaml:"addr,omitempty"`
	Password string `mapstructure:"password" yaml:"password,omitempty"`
	DB       int    `mapstructure:"db" yaml:"db,omitempty"`
}
