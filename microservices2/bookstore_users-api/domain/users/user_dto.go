package users

import (
	"bookstore_users-api/helpers/customerr"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-"`
}

func New() *User {
	return &User{}
}

func (u *User) Validate() *customerr.RestError {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))

	if u.Email == "" {
		return customerr.NewBadRequestErr("invalid email address")
	}

	return nil
}
