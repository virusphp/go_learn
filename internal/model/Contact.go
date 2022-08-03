package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name    string `gorm:"column:name;type:varchar(191)" json:"nama"`
	Alamat  string `gorm:"column:alamat;type:varchar(191)"  json:"alamat"`
	No_telp string `gorm:"column:no_telp;type:varchar(191)"  json:"-"`
}

func (Contact) TableName() string {
	return "contacts"
}
