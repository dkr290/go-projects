package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/helpers/customerr"
)

func GetUser() {

}

func CreateUser(user users.User) (*users.User, *customerr.RestError) {

	return &user, nil
}
func FindUser() {

}
