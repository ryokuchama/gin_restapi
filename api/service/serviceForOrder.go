package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ryokuchama/gin_restapi/api/database"
	"github.com/ryokuchama/gin_restapi/api/entity"
)

// データベース処理
type Service struct{}

type order entity.order

// メニューモデルの作成
func (s Service) CreateOrderModel(c *gin.context) (order, error) {
	db := db.GetOrderDB()
	var o order

	if err := c.BindJSON(&m).Error; err != nil {
		return o, err
	}

	if err := db.Create(&m).Error; err != nil {
		return o, err
	}

	return o, nil
}