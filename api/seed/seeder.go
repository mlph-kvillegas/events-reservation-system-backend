package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mlph-kvillegas/events-reservation-system-backend/api/models"
)

var users = []models.User{
	models.User{
		FirstName:     "Admin",
		LastName:      "Administrator",
		Username:      "admin",
		Email:         "ideyatech_admin@ideyatech.com",
		Password:      "id3y4t3ch",
		ContactNumber: "0910293847561",
		Role:          "ADMIN",
		Status:        true,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Reservation{}, &models.Venue{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Venue{}, &models.Reservation{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
