package api

import (
	"fmt"
	"time"
)

// ProductAPI contains the api methods available for the Product model
type ProductAPI interface {
	// Get provides all products, with the option of listing by user,
	// if an id is passed via header
	Get(userID int64) Response
}

// Response represents the request response
type Response struct {
	Code int
	Path int64
	Body interface{}
}

// ResponseBody represents the standard response body
type ResponseBody struct {
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
	Code    int       `json:"code"`
}

// Error is responsible for encapsulating errors generated by API methods
type Error struct {
	Cause    error
	Response Response
}

func (err *Error) Error() string {
	return fmt.Sprintf("%v %s", err.Response.Code, err.Cause.Error())
}
