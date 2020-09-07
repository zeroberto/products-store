package sqldbdriver

import (
	"database/sql"

	"github.com/zeroberto/products-store/users-data/driver/dbdriver"
)

// SQLDBGenericDriver is responsible for performing operations on a SQL database
type SQLDBGenericDriver struct {
	DB *sql.DB
}

// Exec is responsible for execute an SQL statement
func (driver *SQLDBGenericDriver) Exec(query string, args ...interface{}) (sql.Result, error) {
	return driver.DB.Exec(query, args...)
}

// Prepare a sql statement for the SQL database
func (driver *SQLDBGenericDriver) Prepare(query string, args ...interface{}) (*sql.Stmt, error) {
	return driver.DB.Prepare(query)
}

// PrepareAndExec is responsible for prepare and execute a sql statement for future execution
func (driver *SQLDBGenericDriver) PrepareAndExec(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := driver.Prepare(query)

	defer stmt.Close()

	if err != nil {
		return nil, &dbdriver.Error{Cause: err}
	}

	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, &dbdriver.Error{Cause: err}
	}

	return result, err
}

// Query is responsible for executing an sql command and returning multiple lines
// for the SQL database
func (driver *SQLDBGenericDriver) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return driver.DB.Query(query, args...)
}

// QueryRow is responsible for executing an sql command and returning a single line
// for the SQL database
func (driver *SQLDBGenericDriver) QueryRow(query string, args ...interface{}) *sql.Row {
	return driver.DB.QueryRow(query, args...)
}
