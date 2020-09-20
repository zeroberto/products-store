package dbdriverfactory

import (
	"github.com/zeroberto/products-store/discount-calculator/config"
	"github.com/zeroberto/products-store/discount-calculator/container"
	"github.com/zeroberto/products-store/discount-calculator/container/factory/dbfactory"
	"github.com/zeroberto/products-store/discount-calculator/driver/dbdriver"
	"github.com/zeroberto/products-store/discount-calculator/driver/dbdriver/nosqldbdriver"
	"go.mongodb.org/mongo-driver/mongo"
)

// DBDriverFactory is responsible for providing instances of database driver structs
type DBDriverFactory struct{}

// MakeNoSQLDBDriver is responsible for providing an instance of NoSQLDBDriver
func (dbdf *DBDriverFactory) MakeNoSQLDBDriver(c container.Container, dbc *config.DBConfig) dbdriver.NoSQLDBDriver {
	dbf := &dbfactory.DBFactory{}
	return &nosqldbdriver.MongoDBDriver{
		DB: func() (*mongo.Database, error) {
			return dbf.MakeMongoDB(c, dbc)
		},
	}
}
