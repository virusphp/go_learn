package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name    string `gorm:"column:name;type:varchar(191)" json:"name"`
	Address string `gorm:"column:address;type:varchar(191)"  json:"address"`
	Phone   string `gorm:"column:phone;type:varchar(191)"  json:"phone"`
}

func (Contact) TableName() string {
	return "contacts"
}
