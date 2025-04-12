package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/madhavarora03/restaurant-management-system/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/table", controller.GetTables())
	incomingRoutes.GET("/table/:tableId", controller.GetTable())
	incomingRoutes.POST("/table", controller.CreateTable())
	incomingRoutes.PATCH("/table/:tableId", controller.UpdateTable())
}
