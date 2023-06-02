package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/helpers/customerr"
)

var u = users.New()

func GetUser(userId int64) (*users.User, *customerr.RestError) {

	u.Id = userId
	// result := &users.User{
	// 	Id: userId,
	// }
	if err := u.Get(); err != nil {
		return nil, err
	}
	return u, nil

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
