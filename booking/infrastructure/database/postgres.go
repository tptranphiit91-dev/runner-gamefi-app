package database

import (
	"fmt"
	"sync"
	"booking/domain/entity"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database represents the database connection
// Implements Singleton Pattern
type Database struct {
	DB *gorm.DB
}

var (
	instance *Database
	once     sync.Once
	mu       sync.Mutex
)

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// GetInstance returns the singleton instance of Database
// Singleton Pattern: Ensures only one database connection exists
func GetInstance(config *Config) (*Database, error) {
	var err error
	
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
		)
		
		db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		
		if dbErr != nil {
			err = dbErr
			return
		}
		
		instance = &Database{DB: db}
		
		// Auto migrate tables
		err = instance.AutoMigrate()
	})
	
	return instance, err
}

// AutoMigrate runs database migrations
func (d *Database) AutoMigrate() error {
	return d.DB.AutoMigrate(
		&entity.User{},
	)
}

// Close closes the database connection
func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// ResetInstance resets the singleton instance (useful for testing)
func ResetInstance() {
	mu.Lock()
	defer mu.Unlock()
	instance = nil
	once = sync.Once{}
}

