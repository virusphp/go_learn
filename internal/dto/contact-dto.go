package dto

type ContactDTO struct {
	Name    string `json:"nama" binding:"required,min=3"`
	Address string `json:"alamat" binding:"required"`
	Phone   string `json:"no_telp" binding:"required,min=8"`
}
