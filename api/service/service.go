package service

import (
	"github.com/gin-gonic/gin"
	"./database"
	"./entity"
)

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
func (s Service) UpdateByID(id int, c *gin.Context) (menu, error) {

}