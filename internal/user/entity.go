package user

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required" gorm:"type:varchar(100)"`
	Email    string `json:"email" validate:"required,email" gorm:"type:varchar(100); uniqueIndex; not null"`
	Password string `json:"password" validate:"required,min=6,max=100" gorm:"type:varchar(100) not null"`
}
