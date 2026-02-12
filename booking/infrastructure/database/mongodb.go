package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB represents the MongoDB database connection
// Implements Singleton Pattern
type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var (
	mongoInstance *MongoDB
	mongoOnce     sync.Once
	mongoMu       sync.Mutex
)

// MongoConfig holds MongoDB configuration
type MongoConfig struct {
	URI      string
	Database string
	Timeout  int // Connection timeout in seconds
}

// GetMongoInstance returns the singleton instance of MongoDB
// Singleton Pattern: Ensures only one MongoDB connection exists
func GetMongoInstance(config *MongoConfig) (*MongoDB, error) {
	var err error

	mongoOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Timeout)*time.Second)
		defer cancel()

		// Set client options
		clientOptions := options.Client().ApplyURI(config.URI)

		// Connect to MongoDB
		client, clientErr := mongo.Connect(ctx, clientOptions)
		if clientErr != nil {
			err = clientErr
			return
		}

		// Ping the database to verify connection
		pingErr := client.Ping(ctx, nil)
		if pingErr != nil {
			err = pingErr
			return
		}

		mongoInstance = &MongoDB{
			Client:   client,
			Database: client.Database(config.Database),
		}
	})

	return mongoInstance, err
}

// Close closes the MongoDB connection
func (m *MongoDB) Close() error {
	if m.Client == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return m.Client.Disconnect(ctx)
}

// ResetMongoInstance resets the singleton instance (useful for testing)
func ResetMongoInstance() {
	mongoMu.Lock()
	defer mongoMu.Unlock()
	if mongoInstance != nil {
		mongoInstance.Close()
	}
	mongoInstance = nil
	mongoOnce = sync.Once{}
}

// GetCollection returns a MongoDB collection
func (m *MongoDB) GetCollection(name string) *mongo.Collection {
	return m.Database.Collection(name)
}

// Ping checks if the MongoDB connection is alive
func (m *MongoDB) Ping(ctx context.Context) error {
	return m.Client.Ping(ctx, nil)
}

// BuildMongoURI builds MongoDB connection URI from components
func BuildMongoURI(host, port, username, password string) string {
	if username != "" && password != "" {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}
	return fmt.Sprintf("mongodb://%s:%s", host, port)
}

