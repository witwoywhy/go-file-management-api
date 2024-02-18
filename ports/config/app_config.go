package config

import (
	"strings"

	"github.com/spf13/viper"
)

var Config AppConfig = AppConfig{
	AllowFileExtensions: map[string]bool{},
}

type AppConfig struct {
	AllowFileExtensions map[string]bool
	MaxSizeFile         int64
}

func InitAppConfig() {
	initAllowFileExtensions()
}

func initAllowFileExtensions() {
	s := viper.GetString("app.allow-file-extensions")
	extensions := strings.Split(s, "|")
	for _, extension := range extensions {
		Config.AllowFileExtensions[extension] = true
	}

	Config.MaxSizeFile = viper.GetInt64("app.max-size-file")
}
