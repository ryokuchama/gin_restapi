package database

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ryokuchama/gin_restapi/api/entity"
)

// データベースへの接続とクエリ

var (
	db *gorm.DB
	err error
)

// オーダー用のDB接続処理

func InitOrder() {
	db, err = gorm.Open(
		"mysql",
	"host=0.0.0.0 port=3030 user=gorm dbname=menu password=gorm sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

func GetOrderDB() *gorm.DB {
	return db
}

func CloseOrder() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.order{})
}