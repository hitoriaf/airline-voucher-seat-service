package handlers

import (
	"airline-voucher-seat-service/src/handlers/dto"
	"airline-voucher-seat-service/src/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *VoucherHandler) CheckVoucher(c *gin.Context) {
	var req dto.CheckVoucherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
		return
	}

	filter := repositories.VoucherFilter{
		FlightNumber: &req.FlightNumber,
		FlightDate:   &req.FlightDate,
	}

	vouchers, err := h.voucherRepo.FindByFilter(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to retrieve vouchers",
			Message: err.Error(),
		})
		return
	}

	if len(vouchers) == 0 {
		c.JSON(http.StatusOK, dto.CheckVoucherResponse{Exist: false})
		return
	}

	c.JSON(http.StatusOK, dto.CheckVoucherResponse{Exist: true})
}

