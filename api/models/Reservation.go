package models

import (
	"date"
	"time"
)

type Reservation struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Date      date.Date `gorm:"not null" json:"date"`
	Time      time.Time `gorm:"not null" json:"time"`
	Venue     string    `gorm:"not null" json:"venue"`                    //TODO reference to Venue
	Reservee  string    `gorm:"size:100;not null;" json:"contact_person"` //TODO reference to User
	Status    bool      `gorm:"not null;" json:"status"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
