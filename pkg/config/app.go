package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	stringConnetuion := "host=localhost user=osmait password=123456 dbname=testGO port=5432 sslmode=disable"
	d, err := gorm.Open("postgres", stringConnetuion)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
