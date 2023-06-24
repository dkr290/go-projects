package repository

import "github.com/dkr290/go-projects/goforum-app/goforum/models"

type DatabaseRepo interface {
	InsertPost(newPost models.Post) error
	GetUserById(id int) (models.User, error)
	UpdateUser(u models.User) error
	AuthenticateUser(email, password string) (int, string, error)
}
