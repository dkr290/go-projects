package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the db is completed")
	}

	for i := 0; i <= 10; i++ {
		conn := db.Ping()
		if i > 10 {
			fmt.Println("Failed to connect to the database")
			db.Close()
			return
		}
		if conn != nil {
			fmt.Printf("Trying to connect to the database %d time \n", i)
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("Connected to the database")
			break
		}

	}
	defer db.Close()

	err = createTables(db)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Create tables")
	}

	// err = insertSopmeData(db)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err = deleteFunc(db)
	if err != nil {
		log.Fatal(err)
	}

	err = queryDB(db)
	if err != nil {
		log.Fatal(err)
	}

	err = queryRow(db)
	if err != nil {
		log.Fatal(err)
	}

	err = updataFunc(db)
	if err != nil {
		log.Fatal(err)
	}
}

func createTables(db *sql.DB) error {
	DBCreate := `
	CREATE TABLE IF NOT EXISTS Number (
	  Number integer NOT NULL,
	  Property text  COLLATE pg_catalog."default" NOT NULL
	)
	WITH (
	   OIDS = FALSE
	)
  TABLESPACE pg_default;
	ALTER TABLE Number
	OWNER to postgres;
	`
	_, err := db.Exec(DBCreate)
	if err != nil {
		return errors.New("cannot create the tables" + err.Error())
	}
	return nil
}

func insertSopmeData(db *sql.DB) error {
	insert, err := db.Prepare("INSERT INTO Number VALUES($1,$2)")
	if err != nil {
		return errors.New("cannot insert the data " + err.Error())
	}
	var prop string
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			prop = "Even"
		} else {
			prop = "Odd"
		}
		_, err = insert.Exec(i, prop)
		if err != nil {
			return err
		} else {
			fmt.Println("The number:", i, "is", prop)
		}
	}
	defer insert.Close()
	return nil
}

func queryDB(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM number")
	if err != nil {
		return errors.New("cannot select from the db " + err.Error())
	}
	var number, property string
	for rows.Next() {
		err := rows.Scan(&number, &property)
		if err != nil {
			return err
		}
		fmt.Printf("Retreived data from db %s %s\n", number, property)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	err = rows.Close()
	if err != nil {
		return err
	}
	return nil
}

func queryRow(db *sql.DB) error {
	var property string
	number := 2
	queryRow, err := db.Prepare("SELECT property FROM number where number=$1")
	if err != nil {
		return err
	}

	err = queryRow.QueryRow(number).Scan(&property)
	if err != nil {
		return err
	}
	fmt.Printf("The property with number %d is %s\n", number, property)
	err = queryRow.Close()
	if err != nil {
		return err
	}
	return nil
}

func updataFunc(db *sql.DB) error {
	update := "UPDATE number SET property =$1 where number = $2"
	s, err := db.Prepare(update)
	if err != nil {
		return fmt.Errorf("an error with prepare update %v", err)
	}

	updateResult, err := s.Exec("Updated something", 3)
	if err != nil {
		return fmt.Errorf("error with update %v", err)
	}
	updRecods, err := updateResult.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("The number of records updated %d \n", updRecods)
	return nil
}

func deleteFunc(db *sql.DB) error {
	delS := "DELETE FROM number where number = $1"
	s, err := db.Prepare(delS)
	if err != nil {
		return fmt.Errorf("error on prepare delete %v", err)
	}

	deleteResult, err := s.Exec(4)
	if err != nil {
		return fmt.Errorf("error while deleting %v", err)
	}
	rowsAffected, err := deleteResult.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("the deleted rows are", rowsAffected)
	return nil
}
