package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("blue-vending-machine.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
