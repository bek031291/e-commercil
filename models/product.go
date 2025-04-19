package models

import "gorm.io/gorm"

//type Product struct {
//	gorm.Model
//	Name        string   `json:"name"`
//	Brand       string   `json:"brand"`
//	Description string   `json:"description"`
//	Price       float64  `json:"price"`
//	CategoryID  uint     `json:"category_id"`
//	Category    Category `gorm:"foreignKey:CategoryID"`
//	// 👇 это поле НЕ сохраняется в БД, только для ответа
//	LikesCount int64 `gorm:"-" json:"likes_count"`
//}

type Product struct {
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"category_id"`
}

type Productold struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Brand       string  `json:"brand"`
	CategoryID  uint    `json:"category_id"`
}
