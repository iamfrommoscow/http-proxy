package main

import (
	"proxy/internal/pkg/helpers"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("auth")
	if err := viper.ReadInConfig(); err != nil {
		helpers.LogMessage(err.Error())
		return
	}
	port = viper.GetString("port")
	err := postgres.Connect()
	if err != nil {
		helpers.LogMessage(err.Error())
		return
	}
	server.StartApp(params)
}
