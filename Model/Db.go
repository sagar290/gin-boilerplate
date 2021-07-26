package model

import (
	"strings"

	"github.com/spf13/viper"
)

var db string

func init() {
	viper.SetEnvPrefix("gin-boilerplate")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/gin-boilerplate/")
	viper.AddConfigPath("$HOME/.gin-boilerplate")
	viper.AddConfigPath(".")
}
