package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"userdata/internal/models"
)

type Database interface{}

type PsqlDatabase struct {
	db *sql.DB
}

func InitDb(config DbConfig, numRetries int) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.DBUser, config.DBPassword, config.DBName)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, fmt.Errorf(" cannot open the database %v", err)
	} else {
		fmt.Println("The connection to the db is completed")
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

	db.Close()
	return db, nil
}

func (p *PsqlDatabase) createTables() error {
	DBCreate := `
	CREATE TABLE IF NOT EXISTS users (
	  Id intteger NOT NULL [primary_key]
	  User string NOT NULL,
	  Email string COLLATE pg_catalog."default" NOT NULL
	)
	WITH (
	   OIDS = FALSE
	)
  TABLESPACE pg_default;
	ALTER TABLE users
	OWNER to postgres;
	`
	_, err := p.db.Exec(DBCreate)
	if err != nil {
		return errors.New("cannot create the tables" + err.Error())
	}
	return nil
}

func (p *PsqlDatabase) InsertUsers(data []models.User) error {
	insert, err := p.db.Prepare("INSERT INTO users VALUES($1,$2,$3)")
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

func (p *PsqlDatabase) GetAllRecords() error
