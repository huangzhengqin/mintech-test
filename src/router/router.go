package router

import (
	"github.com/gin-gonic/gin"
	"handler"
)

func RegisterHandlers () *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.POST("/order/create", handler.CreateOrder)
	v1.POST("/order/info", handler.QueryByOrderId)
	v1.POST("/order/list", handler.Query)
	v1.POST("/order/update", handler.UpdateOrder)

	return router
}
