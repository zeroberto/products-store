package model

import "time"

// User represents the user model of the application
type User struct {
	ID          int64
	DateOfBirth time.Time
}
