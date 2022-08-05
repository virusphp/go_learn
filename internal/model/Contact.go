package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name    string `gorm:"column:name;type:varchar(191)" json:"nama"`
	Address string `gorm:"column:alamat;type:varchar(191)"  json:"alamat"`
	Phone   string `gorm:"column:no_telp;type:varchar(191)"  json:"no_telp"`
}

func (Contact) TableName() string {
	return "contacts"
}
