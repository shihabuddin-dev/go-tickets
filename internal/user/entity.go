package user

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Email    string `json:"email" gorm:"type:varchar(100); uniqueIndex; not null"`
	Password string `json:"password" gorm:"type:varchar(100) not null"`
}
