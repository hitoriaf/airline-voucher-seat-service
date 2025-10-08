package routes

import (
	"airline-voucher-seat-service/src/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine) {
	voucherHandler := handlers.NewVoucherHandler()

	api := router.Group("/api")
	{
		// Voucher routes
		api.POST("/check", voucherHandler.CheckVoucher)       // Check voucher by flight number and date
		api.POST("/generate", voucherHandler.GenerateVoucher) // Generate new voucher
	}
}
