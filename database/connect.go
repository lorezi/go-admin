package database

import (
	"os"

	"github.com/lorezi/go-admin/models"
	"github.com/subosito/gotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	gotenv.Load()
	dsn := os.Getenv("DBUSER") + ":" + os.Getenv("DBPASS") + "@" + os.Getenv("DBHOST") + "/" + os.Getenv("DBNAME") + "?charset=utf8"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = db

	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})

}
