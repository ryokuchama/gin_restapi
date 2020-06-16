package database

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ryokuchama/gin_restapi/api/entity"
)

var (
	db *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open(
		"mysql",
	"host=0.0.0.0 port=3030 user=gorm dbname=menu password=gorm sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.menu{})
}