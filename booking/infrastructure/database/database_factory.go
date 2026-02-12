package database

import (
	"fmt"
	"booking/config"
	"booking/domain/repository"
	"booking/infrastructure/observer"
)

// DatabaseFactory creates database connections and repositories
// Factory Pattern: Creates different database implementations based on type
type DatabaseFactory struct {
	config  *config.Config
	subject *observer.Subject
}

// NewDatabaseFactory creates a new database factory
func NewDatabaseFactory(cfg *config.Config, subject *observer.Subject) *DatabaseFactory {
	return &DatabaseFactory{
		config:  cfg,
		subject: subject,
	}
}

// CreateUserRepository creates a user repository based on database type
func (f *DatabaseFactory) CreateUserRepository() (repository.UserRepository, error) {
	switch f.config.DatabaseType {
	case config.PostgresDB:
		return f.createPostgresUserRepository()
	case config.MongoDB:
		return f.createMongoUserRepository()
	default:
		return nil, fmt.Errorf("unsupported database type: %s", f.config.DatabaseType)
	}
}

// createPostgresUserRepository creates a PostgreSQL user repository
func (f *DatabaseFactory) createPostgresUserRepository() (repository.UserRepository, error) {
	dbConfig := &Config{
		Host:     f.config.Database.Host,
		Port:     f.config.Database.Port,
		User:     f.config.Database.User,
		Password: f.config.Database.Password,
		DBName:   f.config.Database.DBName,
		SSLMode:  f.config.Database.SSLMode,
	}
	
	db, err := GetInstance(dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	
	return NewUserRepository(db.DB, f.subject), nil
}

// createMongoUserRepository creates a MongoDB user repository
func (f *DatabaseFactory) createMongoUserRepository() (repository.UserRepository, error) {
	mongoConfig := &MongoConfig{
		URI:      f.config.Database.MongoURI,
		Database: f.config.Database.MongoDBName,
		Timeout:  f.config.Database.MongoTimeout,
	}
	
	db, err := GetMongoInstance(mongoConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	
	return NewUserRepositoryMongo(db, f.subject), nil
}

// GetDatabaseType returns the current database type
func (f *DatabaseFactory) GetDatabaseType() config.DatabaseType {
	return f.config.DatabaseType
}

// Close closes the database connection based on type
func (f *DatabaseFactory) Close() error {
	switch f.config.DatabaseType {
	case config.PostgresDB:
		dbConfig := &Config{
			Host:     f.config.Database.Host,
			Port:     f.config.Database.Port,
			User:     f.config.Database.User,
			Password: f.config.Database.Password,
			DBName:   f.config.Database.DBName,
			SSLMode:  f.config.Database.SSLMode,
		}
		db, err := GetInstance(dbConfig)
		if err != nil {
			return err
		}
		return db.Close()
	case config.MongoDB:
		mongoConfig := &MongoConfig{
			URI:      f.config.Database.MongoURI,
			Database: f.config.Database.MongoDBName,
			Timeout:  f.config.Database.MongoTimeout,
		}
		db, err := GetMongoInstance(mongoConfig)
		if err != nil {
			return err
		}
		return db.Close()
	default:
		return nil
	}
}

