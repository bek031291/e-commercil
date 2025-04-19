package models

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
}
