package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/dkr290/go-projects/goforum-app/goforum/models"
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
