package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name    string `gorm:"column:name;type:varchar(191)"`
	Alamat  string `gorm:"column:alamat;type:varchar(191)"`
	No_telp string `gorm:"column:no_telp;type:varchar(191)"`
}

func (Contact) TableName() string {
	return "contacts"
}
