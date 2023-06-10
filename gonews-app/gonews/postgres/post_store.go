package postgres

import (
	"fmt"

	"github.com/dkr290/go-projects/gonews"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostStore struct {
	*sqlx.DB
}

func (p *PostStore) Post(id uuid.UUID) (gonews.Post, error) {
	var post gonews.Post
	if err := p.Get(&post, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return gonews.Post{}, fmt.Errorf("error selecting the post %w", err)
	}
	return post, nil
}

func (p *PostStore) PostsByThread(threadID uuid.UUID) ([]gonews.Post, error) {
	var posts []gonews.Post
	var query = `SELECT
	               posts.*,
	               COUNT(comments.*) AS comments_count
                   FROM posts
                   LEFT JOIN comments ON comments.post_id = posts.id
                   WHERE thread_id = $1
                   GROUP BY posts.id
                   ORDER BY votes DESC`
	if err := p.Select(&posts, query, threadID); err != nil {
		return []gonews.Post{}, fmt.Errorf("error getting many threads in select all %w", err)
	}
	return posts, nil
}

func (p *PostStore) Posts() ([]gonews.Post, error) {
	var posts []gonews.Post
	var query = `SELECT
	               posts.*,
	               COUNT(comments.*) AS comments_count,
				   threads.title AS thread_title
                   FROM posts
                   LEFT JOIN comments ON comments.post_id = posts.id
				   LEFT JOIN threads ON threads.id = posts.thread_id
                   GROUP BY posts.id, threads.title
                   ORDER BY votes DESC`
	if err := p.Select(&posts, query); err != nil {
		return []gonews.Post{}, fmt.Errorf("error getting posts %w", err)
	}

	return posts, nil
}

func (p *PostStore) CreatePost(t *gonews.Post) error {

	if err := p.Get(t, `INSERT INTO posts VALUES($1 ,$2 ,$3, $4,$5) RETURNING *`,
		t.ID,
		t.ThreadID,
		t.Title,
		t.Content,
		t.Votes); err != nil {
		return fmt.Errorf("error inseartiung post %w", err)
	}
	return nil
}

func (p *PostStore) UpdatePost(t *gonews.Post) error {
	if err := p.Get(t, `UPDATE posts SET thread_id = $1, title = $2, content = $3, votes = $4 WHERE id = $5 RETURNING *`,
		t.ThreadID,
		t.Title,
		t.Content,
		t.Votes,
		t.ID); err != nil {
		return fmt.Errorf("error updating post %w", err)
	}
	return nil
}

func (p *PostStore) DeletePost(id uuid.UUID) error {
	if _, err := p.Exec(`DELETE FROM posts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting post %w", err)
	}
	return nil
}
