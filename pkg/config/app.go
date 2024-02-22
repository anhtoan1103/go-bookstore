package config

import (
	_ "github.com/jinnzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	connectString := "akhil:Axlesharmal@12@/simplerest?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", connectString)

	if err != nil {
		panic(err)
	}

	db = d
}

// that will help use the db in others files.
func GetDB() *gorm.DB {
	return db
}
