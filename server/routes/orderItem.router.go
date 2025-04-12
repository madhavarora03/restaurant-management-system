package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/madhavarora03/restaurant-management-system/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItem", controller.GetOrderItems())
	incomingRoutes.GET("/orderItem/:orderItemId", controller.GetOrderItem())
	incomingRoutes.GET("/orderItem/order/:orderId", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItem", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItem/:orderItemId", controller.UpdateOrderItem())
}
