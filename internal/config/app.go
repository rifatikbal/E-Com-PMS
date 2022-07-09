package config

import (
	"github.com/spf13/viper"
)

var app *AppCfg

type AppCfg struct {
	Host string
	Port int
}

func LoadAppCfg() {
	app = &AppCfg{
		Host: viper.GetString("app.host"),
		Port: viper.GetInt("app.port"),
	}
}

func App() *AppCfg {
	return app
}
