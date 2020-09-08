package driverfactory

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/zeroberto/products-store/users-data/config"
	"github.com/zeroberto/products-store/users-data/container"
	"github.com/zeroberto/products-store/users-data/driver/dbdriver"
	"github.com/zeroberto/products-store/users-data/driver/dbdriver/sqldbdriver"
)

// DBDriverFactory is responsible for providing instances of database driver structs
type DBDriverFactory struct{}

// MakeSQLDBDriver is responsible for providing an instance of SQLDBDriver
func (dsf *DBDriverFactory) MakeSQLDBDriver(c container.Container, dbc *config.DBConfig) (dbdriver.SQLDBDriver, error) {
	db, ok := c.Get(container.SQLConnKey)
	if ok {
		db := db.(*sql.DB)
		sdbd := &sqldbdriver.SQLDBGenericDriver{DB: db}
		return sdbd, nil
	}

	dsn := createSQLDataSourceName(dbc)

	db, err := sql.Open(dbc.Type, dsn)
	if err != nil {
		return nil, err
	}

	err = db.(*sql.DB).Ping()
	if err != nil {
		return nil, err
	}

	sdbd := &sqldbdriver.SQLDBGenericDriver{DB: db.(*sql.DB)}
	c.Put(container.SQLConnKey, sdbd)

	return sdbd, nil
}

func createSQLDataSourceName(dbc *config.DBConfig) string {
	user, pass := getSQLCredentials(&dbc.AuthConfig)

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?charset=utf8", user, pass, dbc.Host, dbc.Port, dbc.Database)
}

func getSQLCredentials(ac *config.AuthConfig) (user string, pass string) {
	if ac.Type == "env" {
		user = os.Getenv(ac.User)
		pass = os.Getenv(ac.Pass)
	} else {
		user = "t"
		pass = "s"
	}
	return
}
