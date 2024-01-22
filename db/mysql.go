package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(uri string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxIdleTime(time.Hour)

	return db, nil
}
