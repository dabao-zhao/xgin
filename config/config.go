package config

import (
	"bytes"
	"path"
	"strings"

	"github.com/spf13/viper"
)

func Load(file string, v interface{}) error {
	fileType := strings.ToLower(path.Ext(file))

	viper.SetConfigType(fileType)

	r := bytes.NewReader([]byte(file))
	if err := viper.ReadConfig(r); err != nil {
		return err
	}
	if err := viper.Unmarshal(&v); err != nil {
		return err
	}

	return nil
}
