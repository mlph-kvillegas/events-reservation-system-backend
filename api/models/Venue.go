package models

import (
	"time"
)

type Venue struct {
	ID            uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name          string    `gorm:"size:255;not null;unique" json:"name"`
	Location      string    `gorm:"size:255;not null;" json:"location"`
	Capacity      int       `gorm:"not null" json:"capacity"`
	ContactPerson User      `json:"contact_person"`
	Status        bool      `gorm:"not null;" json:"status"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
