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

func UpdateUser(isPartial bool, user users.User) (*users.User, *customerr.RestError) {

	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email

	}

	if err := current.Update(); err != nil {

		return nil, err
	}

	return current, nil

}

func DeleteUser(userID int64) *customerr.RestError {
	currUser := &users.User{
		Id: userID,
	}

	return currUser.Delete()
}
func FindUser() {

}
