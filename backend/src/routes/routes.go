package routes

import (
	"airline-voucher-seat-service/src/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	voucherHandler := handlers.NewVoucherHandler()

	api := router.Group("/api")
	{
		api.POST("/check", voucherHandler.CheckVoucher)
		api.POST("/generate", voucherHandler.GenerateVoucher)
	}
}
