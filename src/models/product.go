package models

type Category struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Product struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
}
