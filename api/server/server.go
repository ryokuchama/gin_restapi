package server

import (
	"github.com/gin-gonic/gin"
    "github.com/ryokuchama/gin_restapi/api/controller"
)

//初期化
func Init() {
	r := router()
	r.RUN()
}

// webサーバ起動とルーティング
func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		// メニュー関係
		ctrl := controller.Controller{}
		u.GET("", ctrl.Index)
		u.POST("/", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.Delete("/:id", ctrl.Delete)

		// オーダー関係
		forOrder := controller.Controller{}
		u.POST("/order", forOrder.RegisterOrder)
	}

	return r
}