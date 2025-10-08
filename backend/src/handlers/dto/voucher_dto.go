package dto

type CheckVoucherRequest struct {
	FlightNumber string `json:"flightNumber" binding:"required"`
	FlightDate   string `json:"date" binding:"required"`
}

type GenerateVoucherRequest struct {
	CrewName     string `json:"name" binding:"required"`
	CrewID       string `json:"id" binding:"required"`
	FlightNumber string `json:"flightNumber" binding:"required"`
	FlightDate   string `json:"date" binding:"required"`
	AircraftType string `json:"aircraft" binding:"required"`
}

type GenerateVoucherResponse struct {
	Success bool     `json:"success"`
	Seats   []string `json:"seats"`
}

type CheckVoucherResponse struct {
	Exist bool `json:"exist"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
