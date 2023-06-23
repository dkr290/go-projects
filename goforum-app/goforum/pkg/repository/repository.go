package repository

import "github.com/dkr290/go-projects/goforum-app/goforum/models"

type DatabaseRepo interface {
	InsertPost(newPost models.Post) error
}
