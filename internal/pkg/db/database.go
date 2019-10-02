package db

import (
	sq "database/sql"
	"io/ioutil"
	"proxy/internal/pkg/helpers"
	"strconv"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var connection *sq.DB = nil

var connectionConfig pgx.ConnConfig
var connectionPoolConfig = pgx.ConnPoolConfig{
	MaxConnections: 8,
}

func initDB() error {
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")
	viper.SetConfigName("postgres")
	if err := viper.ReadInConfig(); err != nil {
		helpers.LogMessage("Can't find db config: " + err.Error())
		return err
	}
	connectionConfig = pgx.ConnConfig{
		Host:     viper.GetString("db.host"),
		Port:     uint16(viper.GetInt("db.port")),
		Database: viper.GetString("db.database"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
	}
	psqlURI := "postgresql://" + connectionConfig.User
	if len(connectionConfig.Password) > 0 {
		psqlURI += ":" + connectionConfig.Password
	}
	psqlURI += "@" + connectionConfig.Host + ":" + strconv.Itoa(int(connectionConfig.Port)) + "/" + connectionConfig.Database + "?sslmode=disable"
	var err error
	connection, err = sq.Open("postgres", psqlURI)
	if err != nil {
		return err
	}

	return nil
}

func Connect() error {
	if connection != nil {
		return nil
	}
	err := initDB()
	if query, err := ioutil.ReadFile("internal/pkg/db/init.sql"); err != nil {
		helpers.LogMessage(err.Error())
		return err
	} else {
		if _, err := Exec(string(query)); err != nil {
			helpers.LogMessage(err.Error())
			return err
		}
	}
	if err != nil {
		helpers.LogMessage("Can't connect to db: " + err.Error())
		return err
	}
	helpers.LogMessage("Connected to db")
	return nil
}
