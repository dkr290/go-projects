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
	queryGetUser    = "SELECT id,first_name,last_name,email,date_created from users WHERE id=?"
)

var dbClient = usersdatabase.New()

func (user *User) Get() *customerr.RestError {

	if err := dbClient.Ping(); err != nil {
		panic(err)
	}

	stmt, err := dbClient.Prepare(queryGetUser)
	if err != nil {
		return customerr.NewInternalServerError(err.Error())

	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			fmt.Println(err)
			return customerr.NewNotFoundErr(
				fmt.Sprintf("user with id %d not found", user.Id))
		}

		return customerr.NewInternalServerError(
			fmt.Sprintf("error when trying to get user %d  %s", user.Id, err.Error()))
	}

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
