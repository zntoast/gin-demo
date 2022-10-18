package config

type Mysql struct {
	DbName   string `mapstructure:"dbName" yaml:"dbName,omitempty"`
	IP       string `mapstructure:"ip" yaml:"ip,omitempty"`
	Port     string `mapstructure:"port" yaml:"port,omitempty"`
	UserName string `mapstructure:"userName" yaml:"userName,omitempty"`
	Password string `mapstructure:"password" yaml:"password,omitempty"`
	Config   string `mapstructure:"config" yaml:"config,omitempty"`
}

func (m *Mysql) Dsn() string {
	return m.UserName + ":" + m.Password + "@tcp(" + m.IP + ":" + m.Port + ")/" + m.DbName + "?" + m.Config
}
