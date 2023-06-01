package users

import (
	usersdatabase "bookstore_users-api/datasources/mysql/users_database"
	"bookstore_users-api/helpers/customerr"
	"bookstore_users-api/helpers/datehelpers"
	"fmt"
	"strings"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name,email,date_created) VALUES(?,?,?,?);"
)

var (
	usersDB = make(map[int64]*User)
)
var dbClient = usersdatabase.New()

func (user *User) Get() *customerr.RestError {

	if err := dbClient.Ping(); err != nil {
		panic(err)
	}

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

	stmt, err := dbClient.Prepare(queryInsertUser)
	if err != nil {
		return customerr.NewInternalServerError(err.Error())

	}
	defer stmt.Close()

	user.DateCreated = datehelpers.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return customerr.NewBadRequestErr(fmt.Sprintf("email %s already exists", user.Email))
		}
		return customerr.NewInternalServerError(
			fmt.Sprintf("error trying to save the user  %s", err.Error()))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return customerr.NewInternalServerError(
			fmt.Sprintf("error trying to save the user  %s", err.Error()))
	}

	user.Id = userID
	return nil

}
