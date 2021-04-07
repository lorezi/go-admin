package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbHost = "tcp(localhost:3306)"
	dbName = "mysql"
	dbUser = "root"
	dbPass = "johnwick3"
)

func Connect() {
	dsn := dbUser + ":" + dbPass + "@" + dbHost + "/" + dbName + "?charset=utf8"

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
}
