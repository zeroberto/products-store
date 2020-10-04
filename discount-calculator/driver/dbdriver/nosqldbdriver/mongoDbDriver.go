package nosqldbdriver

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBDriver is responsible for providing a mongo client driver
type MongoDBDriver struct {
	DB func() (*mongo.Database, error)
}

// GetDocByID is responsible for obtaining a single document from the database
// by _id
func (driver *MongoDBDriver) GetDocByID(hexID string, collection string) (interface{}, error) {
	db, err := driver.DB()
	if err != nil {
		return nil, err
	}

	objID, err := primitive.ObjectIDFromHex(hexID)
	if err != nil {
		return nil, err
	}

	var result bson.M
	err = db.Collection(collection).FindOne(context.Background(), bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
