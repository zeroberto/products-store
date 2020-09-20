package dbdriver

// NoSQLDBDriver is responsible for performing operations on a NoSQL database
type NoSQLDBDriver interface {
	// GetDocByID is responsible for obtaining a single document from the database
	// by _id
	GetDocByID(hexID string, collection string) (interface{}, error)
}
