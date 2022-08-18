package repository

import (
	"fmt"
	"go_learn/internal/model"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() (*[]model.User, error)
	FindUserByName(UserName *string, Limit string, Page string, Order string) (*[]model.User, *int64, error)
	InsertUser(b model.User) (*model.User, error)
	UpdateUser(b model.User) (*model.User, error)
	DeleteUser(b model.User) error
	FindUserByID(UserID uint) (*model.User, error)
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
}

type UserConnection struct {
	connection *gorm.DB
}

// IsDuplicateEmail implements UserRepository
func (db *UserConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	panic("unimplemented")
}

// VerifyCredential implements UserRepository
func (db *UserConnection) VerifyCredential(email string, password string) interface{} {
	var user = model.User{}
	res := db.connection.Debug().Where("email =?", email).Take(&user)
	if res.Error != nil {
		fmt.Println("Email tidak di temukan!")
		return false
	}
	byteHash := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return user
}

// DeleteUser implements UserRepository
func (db *UserConnection) DeleteUser(b model.User) error {
	proses := db.connection.Debug().Delete(&b)

	if proses.Error != nil {
		fmt.Println("ada error proses", proses)
		return nil
	}

	db.connection.Find(&b)
	return nil
}

// FindUserByID implements UserRepository
func (db *UserConnection) FindUserByID(UserID uint) (*model.User, error) {
	data := model.User{}
	proses := db.connection.Debug().Where("id=?", UserID).Find(&data)

	if proses.Error != nil {
		fmt.Println("ada pesan error", proses.Error.Error())
		return nil, proses.Error
	}
	return &data, nil

}

// GetUser implements UserRepository
func (db *UserConnection) GetUsers() (*[]model.User, error) {
	data := []model.User{}
	proses := db.connection.Debug().Find(&data)

	if proses.Error != nil {
		fmt.Println("ada error di proses", proses)
		return nil, proses.Error
	}

	return &data, nil
}

func (db *UserConnection) FindUserByName(search *string, limit string, page string, order string) (*[]model.User, *int64, error) {
	limits, _ := strconv.Atoi(limit)
	pages, _ := strconv.Atoi(page)
	var count_ int64
	offset := (pages - 1) * limits
	var users []model.User
	queryBuider := db.connection.Model(&model.User{})
	if search != nil {
		queryBuider = queryBuider.Debug().Where("name LIKE ? ", "%"+*search+"%")
	}
	prosesCount := queryBuider.Debug().Count(&count_)
	if prosesCount.Error != nil {
		fmt.Println("ada error di proses count", prosesCount.Error)
		return &users, &count_, prosesCount.Error
	}
	proses := queryBuider.Debug().Limit(limits).Offset(offset).Order(order).Find(&users)
	if proses.Error != nil {
		fmt.Println("ada error di proses", proses.Error)
		return &users, &count_, proses.Error
	}

	return &users, &count_, nil
}

// InsertUser implements UserRepository
func (db *UserConnection) InsertUser(b model.User) (*model.User, error) {
	bpassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Failed to hash password")
	}
	b.Password = string(bpassword)
	proses := db.connection.Debug().Save(&b)

	if proses.Error != nil {
		fmt.Println("ada error di proses", proses)
		return nil, proses.Error
	}

	db.connection.Find(&b)
	return &b, nil
}

// UpdateUser implements UserRepository
func (db *UserConnection) UpdateUser(b model.User) (*model.User, error) {
	proses := db.connection.Debug().Updates(&b)

	if proses.Error != nil {
		fmt.Println("ada error di proses", proses)
		return nil, proses.Error
	}

	db.connection.Find(&b)
	return &b, nil
}

func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &UserConnection{
		connection: dbConn,
	}
}
