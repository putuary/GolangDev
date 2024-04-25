package models

type Status struct {
	ID   		uint   `gorm:"primary_key" json:"id"`
	StatusName 	string `json:"status_name"`
}
