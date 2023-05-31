package users

import (
	"bookstore_users-api/helpers/customerr"
	"bookstore_users-api/helpers/datehelpers"
	"fmt"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *customerr.RestError {

	result := usersDB[user.Id]
	if result == nil {
		return customerr.NewNotFoundErr(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil

}

func (user *User) Save() *customerr.RestError {

	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return customerr.NewBadRequestErr(fmt.Sprintf("email %s already registered", user.Email))
		}
		return customerr.NewBadRequestErr(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = datehelpers.GetNowString()

	usersDB[user.Id] = user
	return nil

}
