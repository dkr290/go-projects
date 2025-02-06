package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"userdata/internal/models"
)

type Database interface {
	CreateTables() error
	InsertUsers(data []models.User) error
	GetAllRecords() (m []models.User, custerr error)
}

type PsqlDatabase struct {
	Db *sql.DB
}

func InitDb(config DbConfig, numRetries int) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.DBUser, config.DBPassword, config.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf(" cannot open the database %v", err)
	}
	for i := 0; i <= numRetries; i++ {
		conn := db.Ping()
		if i > numRetries {
			db.Close()
			return nil, fmt.Errorf("Failed to connect to the database %v", err)
		}
		if conn != nil {
			fmt.Printf("Trying to connect to the database %d time \n", i)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("Connected to the database")
			break
		}

	}

	return db, nil
}

func (p *PsqlDatabase) CreateTables() error {
	DBCreate := `
	CREATE TABLE IF NOT EXISTS users (
	  Id INTEGER NOT NULL, 
	  Name VARCHAR(50) NOT NULL,
	  Email VARCHAR(50)  NOT NULL
	)
  TABLESPACE pg_default;
	ALTER TABLE users
	OWNER to postgres;
	`
	_, err := p.Db.Exec(DBCreate)
	if err != nil {
		return errors.New("cannot create the tables " + err.Error())
	}

	return nil
}

func (p *PsqlDatabase) InsertUsers(data []models.User) error {
	insert, err := p.Db.Prepare("INSERT INTO users VALUES($1,$2,$3)")
	if err != nil {
		return errors.New("cannot insert the data " + err.Error())
	}
	for _, u := range data {
		_, err = insert.Exec(&u.Id, &u.Name, &u.Email)
		if err != nil {
			return err
		}
	}
	defer insert.Close()
	return nil
}

func (p *PsqlDatabase) GetAllRecords() (m []models.User, custerr error) {
	query := "SELECT * FROM users"
	var data models.User
	rows, err := p.Db.Query(query)
	if err != nil {
		custerr = errors.New("cannot select from the db " + err.Error())
		return
	}

	for rows.Next() {
		err := rows.Scan(&data.Id, &data.Name, &data.Email)
		if err != nil {
			custerr = err
			return
		}
		m = append(m, data)
	}
	return
}
