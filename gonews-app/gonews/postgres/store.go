package postgres

import (
	"fmt"

	"github.com/dkr290/go-projects/gonews"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	gonews.ThreadStore
	gonews.PostStore
	gonews.CommentStore
}

func NewStore(dataSource string) (*Store, error) {

	db, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		return nil, fmt.Errorf("error opening the database %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database %w", err)
	}

	return &Store{
		ThreadStore:  NewThreadStore(db),
		PostStore:    NewPostStore(db),
		CommentStore: NewCommentStore(db),
	}, nil

}
