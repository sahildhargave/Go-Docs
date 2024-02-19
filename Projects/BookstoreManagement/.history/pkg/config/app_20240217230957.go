package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(
		"",
		"username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local",
	)

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
