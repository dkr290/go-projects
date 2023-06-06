package users

import (
	usersdatabase "bookstore_users-api/datasources/mysql/users_database"
	"bookstore_users-api/helpers/customerr"
	"bookstore_users-api/helpers/datehelpers"
	"bookstore_users-api/helpers/mysqlhelpers"
	"fmt"
	"log"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name,email,date_created) VALUES(?,?,?,?);"
	queryGetUser          = "SELECT id,first_name,last_name,email,date_created from users WHERE id=?;"
	QueryUpdateUser       = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?;"
	QueryDeleteUser       = "DELETE from users where id=?;"
	QueryFindUserByStatus = "SELECT id,first_name,last_name,email,date_created,status FROM users WHERE  status=?"
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

func (user *User) Update() *customerr.RestError {

	stmt, err := dbClient.Prepare(QueryUpdateUser)
	if err != nil {
		return customerr.NewInternalServerError(err.Error())

	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		log.Println(err)
		return mysqlhelpers.ParseError(err)

	}

	return nil

}

func (user *User) Delete() *customerr.RestError {
	stmt, err := dbClient.Prepare(QueryDeleteUser)
	if err != nil {
		return customerr.NewInternalServerError(err.Error())

	}
	defer stmt.Close()
	if _, err := stmt.Exec(user.Id); err != nil {
		return mysqlhelpers.ParseError(err)
	}

	return nil

}

func (user *User) FindByStatus(status string) ([]User, *customerr.RestError) {

	stmt, err := dbClient.Prepare(QueryFindUserByStatus)
	if err != nil {
		return nil, customerr.NewInternalServerError(err.Error())

	}

	stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, customerr.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		//var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysqlhelpers.ParseError(err)

		}
		results = append(results, *user)
	}

	if len(results) == 0 {
		return nil, customerr.NewNotFoundErr(fmt.Sprintf("no user matching status %s", status))
	}

	return results, nil

}
