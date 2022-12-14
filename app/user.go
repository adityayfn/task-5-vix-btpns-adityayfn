package app

import "time"

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id" binding:"required"`
	Username string `gorm:"type:varchar(255)" json:"username" binding:"required"`
	Email    string `gorm:"type:varchar(255);UNIQUE" json:"email" binding:"required,email"`
	Password string `gorm:"->;<-; not null" json:"-" binding:"required"`
	Token    string `gorm:"-" json:"token,omitempty"`
	Photos    *[]Photo `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE" json:"photos,omitempty"`	
	CreatedAt time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
