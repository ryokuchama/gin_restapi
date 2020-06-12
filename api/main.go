package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/service_backend/restapi/database"
)

func main() {
	r := gin.Default()
	r.GET("/getAllMenu", func(c *gin.Context) {
		c.String(200, "Hello world")
	})
	
	db.Init()
	r.Run()

	db.Close()
}
