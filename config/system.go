package config

type System struct {
	Addr   string `mapstructure:"addr" yaml:"addr,omitempty"`
	Model  string `mapstructure:"model" yaml:"model,omitempty"`
	Logdir string `mapstructure:"logdir" yaml:"logdir,omitempty"`
}
