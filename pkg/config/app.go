package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// initializing db variable in which it helps other files to interact with db

var (
	db *gorm.DB
)

// Helps to open a connection with datebase (mysq;)
func Connect() {
	database_name := viper.GetString("DATABASE_NAME")
	database_user := viper.GetString("DATABASE_USER")
	database_password := viper.GetString("DATABASE_PASSWORD")
	database_host := viper.GetString("DATABASE_HOST")
	database_port := viper.GetString("DATABASE_PORT")

	connection_link := database_user + ":" + database_password + "@" + "tcp(" + database_host + ":" + database_port + ")/" + database_user

	connection_link = connection_link + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Info(connection_link)
	d, err := gorm.Open(database_name, connection_link)
	if err != nil {
		panic(err)
	}
	db = d
	log.Info("Succesfully established the connection to mysql database")
}

// return db variable

func GetDB() *gorm.DB {
	return db
}
