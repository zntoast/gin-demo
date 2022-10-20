package config

type Server struct {
	Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Mysql   `mapstructure:"mysql" json:"msyql" yaml:"mysql"`
	System  `mapstructure:"system" json:"system" yaml:"system"`
	Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
}
