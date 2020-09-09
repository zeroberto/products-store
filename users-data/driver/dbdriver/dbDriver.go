package dbdriver

import (
	"database/sql"
)

// SQLDBDriver is responsible for performing operations on a SQL database
type SQLDBDriver interface {
	// Exec is responsible for execute an SQL statement
	Exec(query string, args ...interface{}) (sql.Result, error)
	// Prepare a sql statement for future execution
	Prepare(query string, args ...interface{}) (*sql.Stmt, error)
	// PrepareAndExec is responsible for prepare and execute a sql statement for future execution
	PrepareAndExec(query string, args ...interface{}) (sql.Result, error)
	// Query is responsible for executing an sql command and returning multiple lines
	Query(query string, args ...interface{}) (*sql.Rows, error)
	// QueryRow is responsible for executing an sql command and returning a single line
	QueryRow(query string, args ...interface{}) (*sql.Row, error)
}

// Error is responsible for encapsulating errors generated by operations in the database
type Error struct {
	Cause error
}

func (err *Error) Error() string {
	return err.Cause.Error()
}
