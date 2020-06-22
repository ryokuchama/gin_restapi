package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
    "github.com/ryokuchama/gin_restapi/api/service"
)

// コントローラー処理

type Controller struct{}

// 全メニュー取得
func (pc Controller) Index(c *gin.Context) {
	var s service.Service
	p, err := s.GetAllMenu()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	else{
		c.JSON(200, p)
	}
}

// POST処理
func (pc Controller) Create(c *gin.Context) {
	var s service.Service
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	}
	else {
		c.JSON(201, p)
	}
}

// メニュー更新処理
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service

	p, err := s.UpdateById(id, c)

	if err != nil {
		c.AboutWithStatus(400)
		fmt.Println(err)
	}
	else {
		c.JSON(200, p)
	}
}

// メニュー削除
func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service

	if err := s.DeleteById(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	}
	else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}

// IDの最大値を取得
func (pc Controller) GetMaxID(c *gin.Context) {
	
}