package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/madhavarora03/restaurant-management-system/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoice", controller.GetInvoices())
	incomingRoutes.GET("/invoice/:invoiceId", controller.GetInvoice())
	incomingRoutes.POST("/invoice", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoice/:invoiceId", controller.UpdateInvoice())
}
