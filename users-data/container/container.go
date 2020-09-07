package container

import "github.com/zeroberto/products-store/users-data/config"

const (
	// SQLConnKey represents the sql connection key that will be used to
	// obtain the connection instance
	SQLConnKey string = "sql.Conn"
	// TimeStampKey represents the time stamp key that will be used to
	// obtain the TimeStamp instance
	TimeStampKey string = "chrono.TimeStamp"
)

// Container is responsible for managing the application's dependencies
type Container interface {
	// Initialize is responsible for initializing all application structs
	Initialize(configFilename string) error
	// Get is responsible for returning the instance of a struct
	Get(key string) (instance interface{}, ok bool)
	// GetAppConfig is responsible for returning de application configuration properties
	GetAppConfig() *config.AppConfig
	// Put is responsible for adding an instance of a struct to the container
	Put(key string, value interface{})
	// Remove is responsible for remove an instance of a struct to the container
	Remove(key string)
}
