package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection(config *PgConfig) (*gorm.DB, error) {

	DSN := fmt.Sprintf("host=%s port=%s user=%s password=%s  dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil

}
