package driverfactory

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Required database driver

	"github.com/zeroberto/products-store/users-data/config"
	"github.com/zeroberto/products-store/users-data/container"
	"github.com/zeroberto/products-store/users-data/driver/dbdriver"
	"github.com/zeroberto/products-store/users-data/driver/dbdriver/sqldbdriver"
)

// DBDriverFactory is responsible for providing instances of database driver structs
type DBDriverFactory struct{}

// MakeSQLDBDriver is responsible for providing an instance of SQLDBDriver
func (dsf *DBDriverFactory) MakeSQLDBDriver(c container.Container, dbc *config.DBConfig) dbdriver.SQLDBDriver {
	return &sqldbdriver.SQLDBGenericDriver{
		DB: func() (*sql.DB, error) {
			return buildSQLDB(c, dbc)
		},
	}
}

func buildSQLDB(c container.Container, dbc *config.DBConfig) (*sql.DB, error) {
	db, ok := c.Get(container.SQLConnKey)
	if ok {
		if err := db.(*sql.DB).Ping(); err == nil {
			return db.(*sql.DB), nil
		}
		c.Remove(container.SQLConnKey)
	}

	dsn := createSQLDataSourceName(dbc)

	newDb, err := sql.Open(dbc.DatabaseType, dsn)
	if err != nil {
		return nil, err
	}

	err = newDb.Ping()
	if err != nil {
		return nil, err
	}
	c.Put(container.SQLConnKey, newDb)

	return newDb, nil
}

func createSQLDataSourceName(dbc *config.DBConfig) string {
	user, pass := getSQLCredentials(&dbc.AuthConfig)

	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", dbc.DatabaseType, user, pass, dbc.Host, dbc.Port, dbc.Database)
}

func getSQLCredentials(ac *config.AuthConfig) (user string, pass string) {
	if ac.Type == "env" {
		user = os.Getenv(ac.User)
		pass = os.Getenv(ac.Pass)
	} else {
		user = ac.User
		pass = ac.Pass
	}
	return
}
