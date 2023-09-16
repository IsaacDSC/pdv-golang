package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnDatabaseStore() *gorm.DB {
	dsn := "host=localhost user=root password=root dbname=store port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate()
	return db
}
