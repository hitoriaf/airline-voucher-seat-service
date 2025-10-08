package handlers

import (
	"airline-voucher-seat-service/src/database"
	"airline-voucher-seat-service/src/handlers/dto"
	"airline-voucher-seat-service/src/libs"
	"airline-voucher-seat-service/src/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VoucherHandler struct {
	voucherRepo *repositories.VoucherRepository
}

func NewVoucherHandler() *VoucherHandler {
	return &VoucherHandler{
		voucherRepo: repositories.NewVoucherRepository(),
	}
}

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

func (h *VoucherHandler) GenerateVoucher(c *gin.Context) {
	var req dto.GenerateVoucherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
		return
	}
	// Validate Flight Date format
	valid, err := libs.NewLibs().ValidateDate(req.FlightDate)
	if err != nil || !valid {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Invalid date format",
			Message: "Flight date must be in DD-MM-YY format",
		})
		return
	}
	//Check apakah ada voucher dengan flight number yang sama di tanggal yang sama
	vouchers, err := h.voucherRepo.FindByFilter(repositories.VoucherFilter{
		FlightNumber: &req.FlightNumber,
		FlightDate:   &req.FlightDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Voucher Repo Error",
			Message: err.Error(),
		})
		return
	}
	// Voucher sudah ada untuk request ini, kembalikan voucher yang ada
	var seats []string
	if len(vouchers) > 0 {
		seats = []string{vouchers[0].Seat1, vouchers[0].Seat2, vouchers[0].Seat3}
		c.JSON(http.StatusOK, dto.GenerateVoucherResponse{
			Success: false,
			Seats:   seats,
		})
		return
	}
	//Generate seats
	seats = libs.NewLibs().GenerateSeats(req.AircraftType)
	voucher := &database.Voucher{
		CrewName:     req.CrewName,
		CrewID:       req.CrewID,
		FlightNumber: req.FlightNumber,
		FlightDate:   req.FlightDate,
		AircraftType: req.AircraftType,
		Seat1:        seats[0],
		Seat2:        seats[1],
		Seat3:        seats[2],
	}

	if err := h.voucherRepo.Create(voucher); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to create voucher",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.GenerateVoucherResponse{
		Success: true,
		Seats:   []string{voucher.Seat1, voucher.Seat2, voucher.Seat3},
	})
}
