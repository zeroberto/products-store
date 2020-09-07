package sqldatastore

import (
	"database/sql"

	"github.com/zeroberto/products-store/users-data/datastore"
	"github.com/zeroberto/products-store/users-data/driver/dbdriver"
	"github.com/zeroberto/products-store/users-data/model"
)

const (
	// QueryUserInfoByID represents a search query for UserInfo by ID in the base
	QueryUserInfoByID string = `SELECT * FROM user_info WHERE id = ?`
)

// UserInfoDataStoreSQL is responsible for implementing the UserInfoDataStore interface,
// using a relational database
type UserInfoDataStoreSQL struct {
	SQLDriver dbdriver.SQLDBDriver
}

// FindByID is responsible for obtaining a user according to the given identifier
func (uids *UserInfoDataStoreSQL) FindByID(ID int64) (*model.UserInfo, error) {
	rows, err := uids.SQLDriver.Query(QueryUserInfoByID)

	defer rows.Close()

	if err != nil {
		return nil, &datastore.Error{Cause: err}
	}

	return toUserInfo(rows)
}

func toUserInfo(rows *sql.Rows) (*model.UserInfo, error) {
	if rows.Next() {
		userInfo, err := rowsToUserInfo(rows)
		if err != nil {
			return nil, &datastore.Error{Cause: err}
		}
		return userInfo, nil
	}
	return nil, nil
}

func rowsToUserInfo(rows *sql.Rows) (*model.UserInfo, error) {
	var userInfo model.UserInfo
	if err := rows.Scan(
		&userInfo.ID,
		&userInfo.FirstName,
		&userInfo.LastName,
		&userInfo.DateOfBirth,
		&userInfo.CreatedAt,
		&userInfo.UpdatedAt,
		&userInfo.DeactivatedAt,
	); err != nil {
		return nil, &datastore.Error{Cause: err}
	}
	return &userInfo, nil
}
