package repository

import (
	"fmt"
	"go_learn/internal/model"

	"gorm.io/gorm"
)

type ContactRepository interface {
	GetContacts() (*[]model.Contact, error)
	InsertContact(b model.Contact) (*model.Contact, error)
	UpdateContact(b model.Contact) (*model.Contact, error)
	DeleteContact(b model.Contact) error
	FindContactByID(ContactID uint64) (*model.Contact, error)
}

type ContactConnection struct {
	connection *gorm.DB
}

// DeleteContact implements ContactRepository
func (db *ContactConnection) DeleteContact(b model.Contact) error {
	proses := db.connection.Debug().Delete(&b)

	if proses.Error != nil {
		fmt.Println("ada error proses", proses)
		return nil
	}

	db.connection.Find(&b)
	return nil
}

// FindContactByID implements ContactRepository
func (*ContactConnection) FindContactByID(ContactID uint64) (*model.Contact, error) {
	panic("unimplemented")
}

// GetContact implements ContactRepository
func (db *ContactConnection) GetContacts() (*[]model.Contact, error) {
	data := []model.Contact{}
	proses := db.connection.Debug().Find(&data)

	if proses.Error != nil {
		fmt.Println("ada error di proses", proses)
		return nil, proses.Error
	}

	return &data, nil

}

// InsertContact implements ContactRepository
func (*ContactConnection) InsertContact(b model.Contact) (*model.Contact, error) {
	panic("unimplemented")
}

// UpdateContact implements ContactRepository
func (*ContactConnection) UpdateContact(b model.Contact) (*model.Contact, error) {
	panic("unimplemented")
}

func NewContactRepository(dbConn *gorm.DB) ContactRepository {
	return &ContactConnection{
		connection: dbConn,
	}
}
