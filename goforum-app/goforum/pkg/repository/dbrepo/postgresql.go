package dbrepo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/dkr290/go-projects/goforum-app/goforum/models"
	"golang.org/x/crypto/bcrypt"
)

//functions for accessing the database

func (m *PostgresDBRepo) InsertPost(newPost models.Post) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	query := `INSERT INTO posts(title,content,user_id) values($1,$2,$3)`

	_, err := m.DBConn.Exec(ctx, query, newPost.Title, newPost.Content, newPost.UserID)

	if err != nil {
		log.Println("Error inserting to the database", err)
		return err
	}

	return nil
}

func (m *PostgresDBRepo) GetUserById(id int) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT name,email,password,acct_created,last_login,user_type,id FROM users WHERE id = $1`

	row := m.DBConn.QueryRow(ctx, query, id)

	var u models.User
	err := row.Scan(
		&u.Name,
		&u.Email,
		&u.Password,
		&u.AcctCreated,
		&u.LastLogin,
		&u.UserType,
		&u.ID,
	)
	if err != nil {
		log.Println("Error selecting to the database by user id select", err)
		return u, err
	}

	return u, nil

}

func (m *PostgresDBRepo) UpdateUser(u models.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE users SET name=$1, email=$2, last_login=$3, user_type=$4`

	_, err := m.DBConn.Exec(ctx, query, u.Name, u.Email, u.Email, time.Now(), u.UserType)

	if err != nil {
		log.Println("Error updating user to the database", err)
		return err
	}

	return nil

}

func (m *PostgresDBRepo) AuthenticateUser(email, password string) (int, string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id int
	var hashedPW string

	query := `SELECT id,password FROM users where email = $1`

	row := m.DBConn.QueryRow(ctx, query, email)

	err := row.Scan(&id, &hashedPW)

	if err != nil {
		log.Println("Error authenticate user", err)
		return id, "", err
	}

	//compare password provided to thje stored in the database

	err = bcrypt.CompareHashAndPassword([]byte(hashedPW), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("password is incorrect")
	} else if err != nil {
		return 0, "", err

	}

	return id, hashedPW, nil

}

func (m *PostgresDBRepo) GetOneArticle() (int, int, string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id, uID int
	var aTitle, aCcount string

	query := `SELECT id, user_id, title,content FROM posts LIMIT 1`

	row := m.DBConn.QueryRow(ctx, query)

	err := row.Scan(&id, &uID, &aTitle, &aCcount)
	if err != nil {
		return id, uID, "", "", err
	}

	return id, uID, aTitle, aCcount, nil
}

func (m *PostgresDBRepo) GetThreeArticles() (models.ArticleList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var artList models.ArticleList

	rows, err := m.DBConn.Query(ctx, "SELECT id, user_id, title, content FROM posts ORDER BY id DESC LIMIT $1", 3)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, uID int
		var title, content string
		err := rows.Scan(&id, &uID, &title, &content)
		if err != nil {
			panic(err)
		}
		artList.ID = append(artList.ID, id)
		artList.UserID = append(artList.UserID, uID)
		artList.Title = append(artList.Title, title)
		artList.Content = append(artList.Content, content)

	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return artList, nil

}
