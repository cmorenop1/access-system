package db

import (
	"github.com/access-module/api/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	//https://gorm.io/docs/connecting_to_the_database.html#SQLite

	db, err := gorm.Open(sqlite.Open("db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})
	
	return db, nil

}
