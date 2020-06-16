package server

import (
	"github.com/gin-gonic/gin"
    "github.com/ryokuchama/api/controller"
)

//初期化

func Init() {
	r := router()
	r.RUN()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := controller.Controller{}
		u.GET("", ctrl.Index)
		u.POST("/", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.Delete("/:id", ctrl.Delete)
	}

	return r
}