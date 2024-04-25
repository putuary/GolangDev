package models

type Group struct {
	ID       		uint   `gorm:"primary_key" json:"id"`
	GroupName 	 	string `json:"name"`
	Description  	string `json:"description"`
	CreatedAt  		string `json:"created_at"`
}
