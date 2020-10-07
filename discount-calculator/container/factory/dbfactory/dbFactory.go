package dbfactory

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/zeroberto/products-store/discount-calculator/config"
	"github.com/zeroberto/products-store/discount-calculator/container"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBFactory is responsible for providing connections instances
type DBFactory struct{}

// MakeMongoDB is responsible for providing an MongoDB client instance
func (dbf *DBFactory) MakeMongoDB(c container.Container, dbc *config.DBConfig) (*mongo.Database, error) {
	if client, ok := c.Get(container.ProductDBConnKey); ok {
		if err := client.(*mongo.Client).Ping(context.Background(), nil); err == nil {
			return client.(*mongo.Client).Database(dbc.Database), nil
		}
		c.Remove(container.ProductDBConnKey)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 9*time.Second)

	defer cancel()

	connectionURL := createMongoDataSourceName(dbc)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURL))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, err
	}

	c.Put(container.ProductDBConnKey, client)

	return client.Database(dbc.Database), nil
}

func createMongoDataSourceName(dbc *config.DBConfig) string {
	user, pass := getAuthCredentials(&dbc.AuthConfig)

	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?authSource=%s", dbc.DatabaseType, user, pass, dbc.Host, dbc.Port, dbc.Database, dbc.AuthConfig.Repo)
}

func getAuthCredentials(ac *config.AuthConfig) (user string, pass string) {
	if ac.Type == "env" {
		user = os.Getenv(ac.User)
		pass = os.Getenv(ac.Pass)
	} else {
		user = ac.User
		pass = ac.Pass
	}
	return
}
