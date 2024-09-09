package config

type WX struct {
	AppID     string `mapstructure:"appid" json:"appid" yaml:"appid"`             // AppID
	AppSecret string `mapstructure:"appsecret" json:"appsecret" yaml:"appsecret"` // AppSecret
}
