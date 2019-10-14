package config

import (
	"github.com/spf13/viper"
)

// 应用程序配置
type appConfig struct {
	// 应用名称
	Name string
	// 运行模式: debug, release, test
	RunMode string
	// 运行 addr
	Addr string
	// 完整 url
	URL string
	// API 前缀
	APIPrefix string
	// secret key
	Key string
}

func newAppConfig() *appConfig {
	// 默认配置
	viper.SetDefault("APP.NAME", "gin_bbs")
	viper.SetDefault("APP.RUNMODE", "release")
	viper.SetDefault("APP.ADDR", ":8080")
	viper.SetDefault("APP.APIPrefix", "/api")
	viper.SetDefault("APP.KEY", "AYCmJy4cYV1H5kpobYOIOvwgYcghg8A1")

	return &appConfig{
		Name:      viper.GetString("APP.NAME"),
		RunMode:   viper.GetString("APP.RUNMODE"),
		Addr:      viper.GetString("APP.ADDR"),
		URL:       viper.GetString("APP.URL"),
		APIPrefix: viper.GetString("APP.APIPrefix"),
		Key:       viper.GetString("APP.KEY"),
	}
}
