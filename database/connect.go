package database

import (
	"os"

	"github.com/subosito/gotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	gotenv.Load()
	dsn := os.Getenv("DBUSER") + ":" + os.Getenv("DBPASS") + "@" + os.Getenv("DBHOST") + "/" + os.Getenv("DBNAME") + "?charset=utf8"

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}
}
