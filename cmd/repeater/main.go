package main

import (
	"fmt"
	"net/http"
	"proxy/internal/pkg/db"
	"proxy/internal/pkg/handlers"
	"proxy/internal/pkg/helpers"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	db.Connect()
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")
	viper.SetConfigName("repeater")
	if err := viper.ReadInConfig(); err != nil {
		helpers.LogMessage("Impossible to read in server config")
	}
	port := viper.GetString("port")
	router := mux.NewRouter()
	router.HandleFunc("/{id:[0-9]+}", handlers.RepeatRequest).Methods("GET", "OPTIONS")
	fmt.Println("started: " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		helpers.LogMessage("Server couldn't start" + err.Error())
	}
}
