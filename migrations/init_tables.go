package migrations

import (
	"chat-messaging/models"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Status{},
		&models.Group{},
		&models.Message{},
		&models.MessageGroup{},
		&models.GroupMember{},
		&models.Attachment{},
	)
}
