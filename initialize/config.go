package initialize

import (
	"fmt"
	"gindemo/global"

	"github.com/spf13/viper"
)

// 初始化配置文件
func InitConfigFile(path string) {
	if path == "" {
		path = "config.yaml"
	}
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic("load config file fail")
	}
	if err := v.Unmarshal(&global.GVA_SERVER); err != nil {
		panic("文件解析失败")
	}
	fmt.Printf("global.GVA_SERVER.Captcha: %v\n", global.GVA_SERVER.Captcha)
}
