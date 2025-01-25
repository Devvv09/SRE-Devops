package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB
var err error

func ConnectDB() error {
	
	LoadEnviromentVariable()
	// fmt.Println(viper.GetInt("DB_PORT"))
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",viper.GetString("DB_HOST"),viper.GetString("DB_PORT"),viper.GetString("DB_USER"),viper.GetString("DB_PASSWORD"),viper.GetString("DB_NAME"))
	fmt.Println(connStr)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return  err
	}
	fmt.Println("Successfully connected to PostgreSQL!")
	return nil
	
}
