package models

type PhotoUpdateModel struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" gorm:"type:varchar(255)" binding:"required"`
	UserID   uint64 `json:"user_id" form:"user_id"`
}

type PhotoCreateModel struct {
	ID       uint64 `json:"id" form:"id"`
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" gorm:"type:varchar(255)" binding:"required"`
	UserID   uint64 `json:"user_id" form:"user_id"`
}
