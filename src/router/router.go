package router

import (
	"github.com/gin-gonic/gin"
	"handler"
)

func RegisterHandlers () *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.POST("/order", handler.CreateOrder)
	v1.GET("/order/:order_id", handler.QueryByOrderId)
	v1.PUT("/order", handler.UpdateOrder)
	v1.GET("/order", handler.Query)
	v1.POST("/upload", handler.Upload)


	return router
}
