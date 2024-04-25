package models

type GroupMember struct {
	UserID 		uint   `json:"user_id"`
	GroupID		uint   `json:"group_id"`
	RoleID 		uint   `json:"role_id"`
	JoinedAt 	string `json:"joined_at"`

}
