package repository

import (
	"fmt"
	"go_learn/internal/model"

	"gorm.io/gorm"
)

type ContactRepository interface {
	AllContact() (*[]model.Contact, error)
	InsertContact(b model.Contact) (*model.Contact, error)
	UpdateContact(b model.Contact) (*model.Contact, error)
	DeleteContact(b model.Contact) error
	FindContactByID(ContactID uint64) (*model.Contact, error)
}

type ContactConnection struct {
	connection *gorm.DB
}

//NewContactRepository creates an instance ContactRepository
func NewContactRepository(dbConn *gorm.DB) ContactRepository {
	return &ContactConnection{
		connection: dbConn,
	}
}
func (db *ContactConnection) AllContact() (*[]model.Contact, error) {
	data := []model.Contact{}
	proses := db.connection.Debug().Find(&data)

	if proses.Error != nil {
		fmt.Println("ada error di proses", proses)
		return nil, proses.Error
	}

	return &data, nil
}
func (db *ContactConnection) InsertContact(b model.Contact) (*model.Contact, error) {

	proses := db.connection.Debug().Save(&b)

	if proses.Error != nil {
		fmt.Println("ada error di proses", proses)
		return nil, proses.Error
	}

	db.connection.Find(&b)
	return &b, nil
}
func (db *ContactConnection) UpdateContact(b model.Contact) (*model.Contact, error) {
	proses := db.connection.Debug().Updates(&b)

	if proses.Error != nil {
		fmt.Println("ada error di proses", proses)
		return nil, proses.Error
	}

	db.connection.Find(&b)
	return &b, nil
}
func (db *ContactConnection) DeleteContact(b model.Contact) error {
	proses := db.connection.Debug().Delete(&b)

	if proses.Error != nil {
		fmt.Println("ada error di proses", proses)
		return nil
	}

	db.connection.Find(&b)
	return nil
}
func (db *ContactConnection) FindContactByID(ContactID uint64) (*model.Contact, error) {
	var model model.Contact
	proses := db.connection.Find(&model)

	if proses.Error != nil {
		fmt.Println("ada error di proses", proses)
		return nil, proses.Error
	}

	return &model, nil
}
