package models

type User struct {
	ID       		uint   `gorm:"primary_key" json:"id"`
	Name 	 		string `json:"name"`
	PhoneNumber    	string `json:"phone_number"`
	Password 		string `json:"password"`
	BirthDate		string `json:"birth_date"`
	// Tambahkan kolom-kolom lain sesuai kebutuhan
}
