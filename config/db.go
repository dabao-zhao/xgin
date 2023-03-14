package config

import (
	"fmt"
	"time"
)

type DbConfig struct {
	Addr            string        `toml:"addr"`
	User            string        `toml:"user"`
	Pass            string        `toml:"pass"`
	DbName          string        `toml:"dbName"`
	Charset         string        `toml:"charset"`
	MaxOpenConn     int           `toml:"maxOpenConn"`
	MaxIdleConn     int           `toml:"maxIdleConn"`
	ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
}

func (d DbConfig) Dsn() string {
	return d.DsnFromDbName(d.DbName)
}

func (d DbConfig) DsnFromDbName(dbName string) string {
	format := "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"
	return fmt.Sprintf(format, d.User, d.Pass, d.Addr, dbName)
}
