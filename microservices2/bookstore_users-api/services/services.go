package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/helpers/customerr"
)

func GetUser(userId int64) (*users.User, *customerr.RestError) {

	result := &users.User{
		Id: userId,
	}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func CreateUser(user users.User) (*users.User, *customerr.RestError) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
func FindUser() {

}
