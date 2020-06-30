package models

import (
	"time"
)

type User struct {
	ID            uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName     string    `gorm:"size:255;not null;" json:"first_name"`
	LastName      string    `gorm:"size:255;not null;" json:"last_name"`
	Email         string    `gorm:"size:100;not null;unique" json:"email"`
	Password      string    `gorm:"size:100;not null;" json:"password"`
	ContactNumber string    `gorm:"size:255;not null;" json:"contact_number"`
	Role          string    `gorm:"size:255;not null;" json:"role"`
	Status        bool      `gorm:"not null;" json:"status"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}