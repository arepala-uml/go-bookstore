package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
)

// initializing db variable in which it helps other files to interact with db

var (
	db *gorm.DB
)

// Host: sql12.freesqldatabase.com
// Database name: sql12735030
// Database user: sql12735030
// Database password: AuLnnFYxhu
// Port number: 3306

// Helps to open a connection with datebase (mysq;)
func Connect() {

	d, err := gorm.Open("mysql", "sql12735030:AuLnnFYxhu@tcp(sql12.freesqldatabase.com:3306)/sql12735030?charset=utf8mb4&parseTime=True&loc=Local")
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
