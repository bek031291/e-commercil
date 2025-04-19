package models

type Category struct {
	//gorm.Model
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}
