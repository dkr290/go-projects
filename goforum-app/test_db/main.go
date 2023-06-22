package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	//connect to the db

	count := 0
	var conn *pgx.Conn
	var err error
	connString := "postgres://postgres:password@localhost:5432/blog_db"

	for {

		//conn, err = sql.Open("pgx", "host=postgres port=5432 dbname=blog_db user=postgres password=password")
		conn, err = pgx.Connect(context.Background(), connString)
		if err != nil && count <= 10 {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			count++
			time.Sleep(10 * time.Second)
		} else if err == nil {
			log.Println("Connected to the database")
			break
		} else if count > 10 {
			log.Fatalln("Too many retries to connect to the database")
			break
		}

	}

	defer conn.Close(context.Background())

	// err = getAllRowData(conn)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = inNewUser(conn, "Peter", "pkeler@abv.bg", "password1234")
	// if err != nil {

	// 	log.Fatal("Cannot insert into the database", err)
	// }

	// err = getUserData(conn, 2)
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	err = updateUserEmail(conn, "newemail@abv.bg", 2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)

	}

	err = deleteUser(conn, 1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)

	}

	err = getAllRowData(conn)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRowData(conn *pgx.Conn) error {
	rows, err := conn.Query(context.Background(), "select id,name,email from users")
	if err != nil {
		log.Println(err)
		return err
	}

	defer rows.Close()

	var name, email string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Data: ", id, name, email)
	}

	fmt.Println("============================================")

	return nil

}

func inNewUser(conn *pgx.Conn, name, email, password string) error {

	timeStamp := time.Now()

	_, err := conn.Exec(context.Background(), `INSERT INTO users(name,email,password,acct_created,last_login,user_type) values($1,$2,$3,$4,$5,$6)`,
		name, email, password, timeStamp, timeStamp, 3)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil

}

func getUserData(conn *pgx.Conn, id int) error {

	var name, email, password string
	var user_type, sId int
	err := conn.QueryRow(context.Background(), "select id, name, email, password, user_type from users where id=$1", id).Scan(&sId, &name, &email, &password, &user_type)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}

	fmt.Println("ID:", sId)
	fmt.Println("Name: ", name)
	fmt.Println("Email:", email)
	fmt.Println("Password:", password)
	fmt.Println("User type", user_type)

	return nil
}

func updateUserEmail(conn *pgx.Conn, newEmail string, id int) error {

	_, err := conn.Exec(context.Background(), "update users set email=$1 where id=$2", newEmail, id)
	if err != nil {

		return err
	}

	return nil
}

func deleteUser(conn *pgx.Conn, id int) error {
	_, err := conn.Exec(context.Background(), "delete from users where id=$1", id)

	if err != nil {
		return err
	}
	return nil
}
