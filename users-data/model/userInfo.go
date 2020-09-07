package model

import "time"

// UserInfo represents the data model that contains a user's main data
type UserInfo struct {
	ID            int64
	FirstName     string
	LastName      string
	DateOfBirth   time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeactivatedAt time.Time
}
