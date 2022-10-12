package models

import "time"

type RegisterModel struct {
	Username  string    `json:"username" form:"username" binding:"required" validate:"min:6"`
	Email     string    `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password  string    `json:"password" form:"password" binding:"required" validate:"min:6"`
	CreatedAt time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
UpdatedAt time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
