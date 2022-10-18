package config

type Captcha struct {
	ImgHeight int `mapstructure:"img-height" yaml:"img-height,omitempty"` // 图像高度
	ImgWidth  int `mapstructure:"img-width" yaml:"img-width,omitempty"`   // 图像宽度
	Long      int `mapstructure:"long" yaml:"long,omitempty"`             // 验证码长度
}
