package datastore

import (
	"github.com/zeroberto/products-store/users-data/model"
)

// UserInfoDataStore represents the user data access interface
type UserInfoDataStore interface {
	// FindByID is responsible for obtaining a user according to the given identifier
	FindByID(ID int64) (*model.UserInfo, error)
}

// Error is responsible for encapsulating errors generated by operations in the data access layer
type Error struct {
	Cause error
}

func (err *Error) Error() string {
	return err.Cause.Error()
}