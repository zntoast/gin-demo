package initialize

import (
	"gindemo/global"

	"github.com/spf13/viper"
)

// 加载配置文件
func InitConfigFile(path string) {
	if path == "" {
		path = "config.yaml"
	}
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic("加载配置文件失败")
	}
	if err := v.Unmarshal(&global.GVA_SERVER); err != nil {
		panic("文件解析失败")
	}
}
