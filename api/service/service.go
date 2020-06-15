package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ryokuchama/gin_restapi/database"
	"github.com/ryokuchama/gin_restapi/entity"
)

// データベース処理

type Service struct{}

type menu entity.menu

// 全メニュー取得
func (s Service) GetAllMenu() ([]menu, error) {
	db := db.GetDB()
	var m []menu

	if err := db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// メニューモデルの作成
func (s Service) CreateModel(c *gin.context) (menu, error) {
	db := db.GetDB()
	var m menu

	if err := c.BindJSON(&m).Error; err != nil {
		return m, err
	}

	if err := db.Create(&m).Error; err != nil {
		return m, err
	}

	return m, nil
}

// メニュー更新
func (s Service) UpdateByID(id string, c *gin.Context) (menu, error) {
	db := db.GetDB()
	var m menu

	if err := db.Where("id = ?", id).First(&m).Error; err != nil {
		return m, err
	}
	if err := c.BindJSON(&m); err != nil {
		return m, err
	}

	db.Save(&m)

	return m, nil
}

// メニュー削除
func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var m menu

	if err := db.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}