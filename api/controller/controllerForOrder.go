package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
    "github.com/ryokuchama/gin_restapi/api/service"
)

// コントローラー処理

type Controller struct{}

// POST処理
func (pc Controller) RegisterOrder(c *gin.Context) {
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