package postgres

import (
	"fmt"

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
	var comment gonews.Comment
	if err := c.DB.Get(&comment, `SELECT * FROM comments WHERE id=$1`, id); err != nil {
		return gonews.Comment{}, fmt.Errorf("error select comment %w", err)
	}
	return comment, nil
}

func (c *CommentStore) CommentsByPost(postID uuid.UUID) ([]gonews.Comment, error) {
	var comments []gonews.Comment
	if err := c.DB.Select(&comments, `SELECT FROM comments WHERE post_id = $1`, postID); err != nil {
		return []gonews.Comment{}, fmt.Errorf("error getting many comments by post_id %w", err)
	}
	return comments, nil
}

func (c *CommentStore) CreateComment(t *gonews.Comment) error {
	if err := c.DB.Get(t, `INSERT INTO comments VALUES($1, $2, $3, $4) RETURNING *`,
		t.ID,
		t.PostID,
		t.Content,
		t.Votes); err != nil {
		return fmt.Errorf("error inserting a comment %w", err)
	}
	return nil
}

func (c *CommentStore) UpdateComment(t *gonews.Comment) error {
	panic("not implemented") // TODO: Implement
}

func (c *CommentStore) DeleteComment(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}
