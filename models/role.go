package models

type Role struct {
	ID   		uint   `gorm:"primary_key" json:"id"`
	RoleName  	string `json:"role_name"`
}
