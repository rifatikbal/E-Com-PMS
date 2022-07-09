package config

import (
	"time"

	"github.com/spf13/viper"
)

var db *DBCfg

type DBCfg struct {
	Host            string
	Port            int
	User            string
	Pass            string
	Name            string
	ConnMaxLifetime time.Duration
}

func LoadDBCfg() {
	db = &DBCfg{
		Host:            viper.GetString("db.host"),
		Port:            viper.GetInt("db.port"),
		User:            viper.GetString("db.user"),
		Pass:            viper.GetString("db.pass"),
		Name:            viper.GetString("db.name"),
		ConnMaxLifetime: viper.GetDuration("db.conn_max_lifetime") * time.Second,
	}
}

func DB() *DBCfg {
	return db
}
