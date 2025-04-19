package models

type User struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	KeycloakID string `gorm:"uniqueIndex" json:"keycloak_id"`
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	IsActive   bool   `json:"is_active"`
	Password   string `gorm:"-" json:"password"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
