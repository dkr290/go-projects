package users

import (
	usersdatabase "bookstore_users-api/datasources/mysql/users_database"
	"bookstore_users-api/helpers/customerr"
	"bookstore_users-api/helpers/datehelpers"
	"bookstore_users-api/helpers/mysqlhelpers"
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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {

		return mysqlhelpers.ParseError(getErr)
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
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysqlhelpers.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlhelpers.ParseError(saveErr)
	}

	user.Id = userID
	return nil

}
