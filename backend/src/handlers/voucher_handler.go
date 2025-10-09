package handlers

import (
	"airline-voucher-seat-service/src/repositories"
)

type VoucherHandler struct {
	voucherRepo *repositories.VoucherRepository
}

func NewVoucherHandler() *VoucherHandler {
	return &VoucherHandler{
		voucherRepo: repositories.NewVoucherRepository(),
	}
}
