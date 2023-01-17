package postgres

import (
	"github.com/dkr290/go-projects/gonews"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewCommentStore(db *sqlx.DB) *CommentStore {
	return &CommentStore{
		DB: db,
	}
}

type CommentStore struct {
	DB *sqlx.DB
}

func (c *CommentStore) Comment(id uuid.UUID) (gonews.Comment, error) {
	panic("not implemented") // TODO: Implement
}

func (c *CommentStore) CommentsByPost(postID uuid.UUID) ([]gonews.Comment, error) {
	panic("not implemented") // TODO: Implement
}

func (c *CommentStore) CreateComment(t *gonews.Comment) error {
	panic("not implemented") // TODO: Implement
}

func (c *CommentStore) UpdateComment(t *gonews.Comment) error {
	panic("not implemented") // TODO: Implement
}

func (c *CommentStore) DeleteComment(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}
