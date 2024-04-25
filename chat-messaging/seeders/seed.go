package seeders

import (
	"chat-messaging/models"

	"github.com/jinzhu/gorm"
)

func Seed(db *gorm.DB) {

	// Seeder untuk Users
	user1 := models.User{Name: "user1", PhoneNumber:"081234567890", Password: "password1", BirthDate: "1990-01-01"}
	user2 := models.User{Name: "user2", PhoneNumber:"081234567891", Password: "password2", BirthDate: "1991-01-01"}

	db.Create(&user1)
	db.Create(&user2)
}
