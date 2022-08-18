package dto

type ContactDTO struct {
	Name    string `json:"nama" binding:"required,min=3"`
	Address string `json:"alamat" binding:"required"`
	Phone   string `json:"no_telp" binding:"required,min=8"`
	UserID  uint32 `gorm:" not null;" json:"-"`
}

type ListContactDTO struct {
	Search *string `json:"search"`
	Limit  string  `json:"limit"`
	Page   string  `json:"page"`
	Order  string  `json:"order"`
}
