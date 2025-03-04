package models

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Username     string `json:"username" gorm:"unique"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}
