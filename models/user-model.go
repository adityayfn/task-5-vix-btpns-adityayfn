package models

import "time"

type UserUpdateModel struct {
	ID       uint64 `json:"id" form:"id"`
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
}
