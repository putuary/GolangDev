package models

type Attachment struct {
	ID   		uint   `gorm:"primary_key" json:"id"`
	MessageID 	string `json:"message_id"`
	FileName  	string `json:"file_name"`
}
