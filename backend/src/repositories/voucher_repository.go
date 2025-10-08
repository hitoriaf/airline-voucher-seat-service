package repositories

import (
	"errors"

	"airline-voucher-seat-service/src/database"

	"gorm.io/gorm"
)

type VoucherRepository struct {
	db *gorm.DB
}

type VoucherFilter struct {
	CrewName     *string
	CrewID       *string
	FlightNumber *string
	FlightDate   *string
	AircraftType *string
}

func NewVoucherRepository() *VoucherRepository {
	return &VoucherRepository{
		db: database.GetDB(),
	}
}

// FindAll returns all vouchers
func (r *VoucherRepository) FindAll() ([]database.Voucher, error) {
	var vouchers []database.Voucher
	result := r.db.Find(&vouchers)
	if result.Error != nil {
		return nil, result.Error
	}
	return vouchers, nil
}

// FindByFilter returns vouchers based on filter criteria
func (r *VoucherRepository) FindByFilter(filter VoucherFilter) ([]database.Voucher, error) {
	var vouchers []database.Voucher
	query := r.db

	if filter.CrewName != nil && *filter.CrewName != "" {
		query = query.Where("crew_name LIKE ?", "%"+*filter.CrewName+"%")
	}
	if filter.CrewID != nil && *filter.CrewID != "" {
		query = query.Where("crew_id = ?", *filter.CrewID)
	}
	if filter.FlightNumber != nil && *filter.FlightNumber != "" {
		query = query.Where("flight_number = ?", *filter.FlightNumber)
	}
	if filter.FlightDate != nil && *filter.FlightDate != "" {
		query = query.Where("flight_date = ?", *filter.FlightDate)
	}
	if filter.AircraftType != nil && *filter.AircraftType != "" {
		query = query.Where("aircraft_type = ?", *filter.AircraftType)
	}

	result := query.Find(&vouchers)
	if result.Error != nil {
		return nil, result.Error
	}
	return vouchers, nil
}

// Create creates a new voucher
func (r *VoucherRepository) Create(voucher *database.Voucher) error {
	result := r.db.Create(voucher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update updates an existing voucher
func (r *VoucherRepository) Update(id int, voucher *database.Voucher) error {
	// Check if voucher exists
	var existingVoucher database.Voucher
	result := r.db.First(&existingVoucher, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("voucher not found")
		}
		return result.Error
	}

	// Update the voucher
	voucher.ID = id
	result = r.db.Save(voucher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete deletes a voucher by ID
func (r *VoucherRepository) Delete(id int) error {
	// Check if voucher exists
	var voucher database.Voucher
	result := r.db.First(&voucher, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("voucher not found")
		}
		return result.Error
	}

	// Delete the voucher
	result = r.db.Delete(&voucher, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}