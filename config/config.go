package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

func Load(file string, v interface{}) error {
	ext := strings.ToLower(path.Ext(file))
	viper.SetConfigType(ext[1:])

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	r := bytes.NewReader(data)
	if err := viper.ReadConfig(r); err != nil {
		return err
	}
	if err := viper.Unmarshal(&v); err != nil {
		return err
	}

	return nil
}
