package db

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/jinzhu/gorm"
	"qor-test/app/models"
	_ "os"
)

var DB, _ = gorm.Open("sqlite3", "qor-test.db")

func init() {
	AutoMigrate(&models.Product{})
}

func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		DB.AutoMigrate(value)
	}
}
