package models

type MessageGroup struct {
	ID     				uint            `gorm:"primary_key" json:"id"`
	Message				string          `json:"message"`
	SenderID 			uint          	`json:"sender_id"`
	ReceivedGroupID 	uint          	`json:"received_group_id"`
	StatusID   			uint          	`json:"status_id"`
	CreatedAt  			string         	`json:"created_at"`
	UpdatedAt  			string         	`json:"updated_at"`
}
