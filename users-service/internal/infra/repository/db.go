package repository

import (
	"fmt"
	"time"

	"users-service/internal/core/port/repository"

	config "users-service/conf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	*gorm.DB
}

func NewDB() (repository.Database, error) {
	db, err := newDatabase()
	if err != nil {
		return nil, err
	}

	return &database{
		db,
	}, nil
}

func newDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.DB.Host,
		config.DB.User,
		config.DB.Password,
		config.DB.Name,
		config.DB.Port,
		config.DB.SslMode,
		config.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(config.DB.ConnMaxLifetimeInMinute))
	sqlDB.SetMaxOpenConns(config.DB.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.DB.MaxOpenConns)
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	return db, err
}

func (db database) Close() error {
	return db.Close()
}

func (db database) GetDB() *gorm.DB {
	return db.DB
}
