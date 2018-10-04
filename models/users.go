package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	// import only by design of GORM
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// ErrNotFound is returned when a resource cannot be found
	// in the database
	ErrNotFound = errors.New("models: resource not found")
)

// User is the gorm model for users
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

// UserService is the public API for users and their CRUD
type UserService struct {
	db *gorm.DB
}

// NewUserService is used to create a new client for the USerService
func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

// Close is used to end the DB connection of this UserService
func (us *UserService) Close() error {
	return us.db.Close()
}

// ByID looks up a user using the provided ID
// If the user is found error is nil
// IF the user is not found we return an ErrNotFound
// If there is another error we will return that with more information
// about what went wrong, including errors thrown from other packages.
//
// Anything other than ErrNotFound should probably be a 500
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	err := us.db.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// DestructieReset drops the user table and rebuilds it
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}
