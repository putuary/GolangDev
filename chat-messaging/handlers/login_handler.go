package handlers

import (
	"chat-messaging/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		var user models.User
		if err := db.Where("phone_number = ?", input.PhoneNumber).First(&user).Error; err != nil {
			c.JSON(404, gin.H{"message": "User not found"})
			return
		}

		// Pastikan untuk memeriksa kata sandi yang benar di sini

		token, err := CreateToken(user.ID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}

		c.JSON(200, gin.H{"token": token})
	}
}
