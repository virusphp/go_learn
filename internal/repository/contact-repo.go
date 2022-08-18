package repository

import (
	"fmt"
	"go_learn/internal/model"
	"strconv"

	"gorm.io/gorm"
)

type ContactRepository interface {
	GetContacts() (*[]model.Contact, error)
	FindContactByName(ContactName *string, Limit string, Page string, Order string) (*[]model.Contact, *int64, error)
	InsertContact(b model.Contact) (*model.Contact, error)
	UpdateContact(b model.Contact) (*model.Contact, error)
	DeleteContact(b model.Contact) error
	FindContactByID(ContactID uint) (*model.Contact, error)
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
func (db *ContactConnection) FindContactByID(ContactID uint) (*model.Contact, error) {
	data := model.Contact{}
	proses := db.connection.Debug().Where("id=?", ContactID).Find(&data)

	if proses.Error != nil {
		fmt.Println("ada pesan error", proses.Error.Error())
		return nil, proses.Error
	}
	return &data, nil

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

func (db *ContactConnection) FindContactByName(search *string, limit string, page string, order string) (*[]model.Contact, *int64, error) {
	limits, _ := strconv.Atoi(limit)
	pages, _ := strconv.Atoi(page)
	var count_ int64
	offset := (pages - 1) * limits
	var contacts []model.Contact
	queryBuider := db.connection.Model(&model.Contact{})
	if search != nil {
		queryBuider = queryBuider.Debug().Where("name LIKE ? ", "%"+*search+"%")
	}
	prosesCount := queryBuider.Debug().Count(&count_)
	if prosesCount.Error != nil {
		fmt.Println("ada error di proses count", prosesCount.Error)
		return &contacts, &count_, prosesCount.Error
	}
	proses := queryBuider.Debug().Limit(limits).Offset(offset).Order(order).Find(&contacts)
	if proses.Error != nil {
		fmt.Println("ada error di proses", proses.Error)
		return &contacts, &count_, proses.Error
	}

	return &contacts, &count_, nil
}

// InsertContact implements ContactRepository
func (db *ContactConnection) InsertContact(b model.Contact) (*model.Contact, error) {
	proses := db.connection.Debug().Save(&b)

	if proses.Error != nil {
		fmt.Println("ada error di proses", proses)
		return nil, proses.Error
	}

	db.connection.Find(&b)
	return &b, nil
}

// UpdateContact implements ContactRepository
func (db *ContactConnection) UpdateContact(b model.Contact) (*model.Contact, error) {
	proses := db.connection.Debug().Updates(&b)

	if proses.Error != nil {
		fmt.Println("ada error di proses", proses)
		return nil, proses.Error
	}

	db.connection.Find(&b)
	return &b, nil
}

func NewContactRepository(dbConn *gorm.DB) ContactRepository {
	return &ContactConnection{
		connection: dbConn,
	}
}
