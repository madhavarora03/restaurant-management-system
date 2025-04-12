package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/madhavarora03/restaurant-management-system/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order", controller.GetOrders())
	incomingRoutes.GET("/order/:orderId", controller.GetOrder())
	incomingRoutes.POST("/order", controller.CreateOrder())
	incomingRoutes.PATCH("/order/:orderId", controller.UpdateOrder())
}
