package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbName = os.Getenv("DB_NAME")
)

func Connection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to create a connection to database")
	}

	return db
}
