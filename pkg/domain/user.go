package domain

type User struct {
	ID    uint   `json:"id" gorm:"unique;not null"`
	Name  string `json:"name"`
	Email string `json:"Email"`
}
