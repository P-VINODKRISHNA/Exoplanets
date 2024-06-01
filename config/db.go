package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

func DBinit() {
	connStr := viper.GetString("POSTGRES_CONNECTION_STRING")

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
