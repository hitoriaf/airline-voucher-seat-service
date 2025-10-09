package libs

import (
	"fmt"
	"math/rand"
	"time"
)

func NewLibs() *Libs {
	return &Libs{}
}

type Libs struct {}

func (l *Libs) GenerateSeats(aircraftType string) []string {
	var rows int
	var letters []string
	switch aircraftType {
	case "ATR":
		rows = 18
		letters = []string{"A", "C", "D", "F"}
	case "Airbus 320", "Boeing 737 Max":
		rows = 32
		letters = []string{"A", "B", "C", "D", "E", "F"}
	}
	var seats []string
	var seatClaimed = map[string]bool{}
	for len(seats) < 3 {
		row := rand.Intn(rows) + 1
		letter := letters[rand.Intn(len(letters))]
		seat := fmt.Sprintf("%d%s", row, letter)

		//prevent duplicate
		if !seatClaimed[seat] {
			seats = append(seats, seat)
			seatClaimed[seat] = true
		}
	}
	return seats
}

func (l *Libs) ValidateDate(date string) (bool, error) {
	// Validate date format (YYYY-MM-DD) using time.Parse
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false, fmt.Errorf("date must be in YYYY-MM-DD format")
	}
	return true, nil
}