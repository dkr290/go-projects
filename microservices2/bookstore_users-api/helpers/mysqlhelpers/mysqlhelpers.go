package mysqlhelpers

import (
	"bookstore_users-api/helpers/customerr"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var (
	errorRows = "no rows in result set"
)

func ParseError(err error) *customerr.RestError {

	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), errorRows) {
			return customerr.NewNotFoundErr("no record found")
		}

		return customerr.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return customerr.NewBadRequestErr("duplicated key")
	}

	return customerr.NewInternalServerError("error when trying saving the user")
}
