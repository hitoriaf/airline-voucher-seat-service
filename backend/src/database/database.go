package database

import (
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Voucher struct {
	ID           int       `json:"id" db:"id" gorm:"primaryKey;autoIncrement"`
	CrewName     string    `json:"crew_name" db:"crew_name"`
	CrewID       string    `json:"crew_id" db:"crew_id"`
	FlightNumber string    `json:"flight_number" db:"flight_number"`
	FlightDate   string    `json:"flight_date" db:"flight_date"`
	AircraftType string    `json:"aircraft_type" db:"aircraft_type"`
	Seat1        string    `json:"seat1" db:"seat1"`
	Seat2        string    `json:"seat2" db:"seat2"`
	Seat3        string    `json:"seat3" db:"seat3"`
	CreatedAt    time.Time `json:"created_at" db:"created_at" gorm:"autoCreateTime"`
}

var db *gorm.DB

func InitDB() {
	var err error
	sqlPath := os.Getenv("DB_PATH")
	db, err = gorm.Open(sqlite.Open(sqlPath), &gorm.Config{})
	if err != nil {

		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Voucher{})
	if err != nil {
		panic("failed to migrate database")
	}
}

func GetDB() *gorm.DB {
	return db
}
