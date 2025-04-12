package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/madhavarora03/restaurant-management-system/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controller.GetMenus())
	incomingRoutes.GET("/menu/:menuId", controller.GetMenu())
	incomingRoutes.POST("/menu", controller.CreateMenu())
	incomingRoutes.PATCH("/menu/:menuId", controller.UpdateMenu())
}
