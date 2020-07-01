package models

import (
	"time"
)

type Reservation struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	DateFrom  time.Time `gorm:"not null" json:"date_from"` // format/parse before saving/getting (Date and Time)
	DateTo    time.Time `gorm:"not null" json:"date_to"`   // format/parse before saving/getting (Date and Time)
	Venue     Venue     `json:"venue"`
	Reservee  User      `json:"contact_person"`
	Status    bool      `gorm:"not null;" json:"status"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
