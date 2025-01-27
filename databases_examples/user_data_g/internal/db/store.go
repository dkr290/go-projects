package db

import (
	"fmt"
	"time"
	"userdata/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	CreateTables() error
	InsertUsers(data []models.User) error
	GetAllRecords() (m []models.User, custerr error)
}

type PsqlDatabase struct {
	Db *gorm.DB
}

func InitDb(config DbConfig, numRetries int) (db *gorm.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.DBUser, config.DBPassword, config.DBName)

	for i := 0; i <= numRetries; i++ {
		db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

		if i > numRetries {
			return nil, fmt.Errorf("Failed to connect to the database %v", err)
		}
		if err != nil {
			fmt.Printf("Trying to connect to the database %d time \n", i)
			fmt.Println(err)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("Connected to the database")
			break
		}
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (p *PsqlDatabase) InsertUsers(data []models.User) error {
	for _, u := range data {
		result := p.Db.Create(&u)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (p *PsqlDatabase) GetAllRecords() (m []models.User, custerr error) {
	result := p.Db.Find(&m)
	if result.Error != nil {
		return nil, result.Error
	}

	return
}
