package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestSelect(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Error: %s has occurred when opening stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"column"}).AddRow("row")
	mock.ExpectQuery("SELECT column FROM table").WillReturnRows(rows)

	var query string = "SELECT column FROM table"
	count, err := Select(db, query)
	if err != nil {
		t.Errorf("Not expecting an error, but got %s instead", err)
	}

	if count < 1 {
		t.Errorf("`Select(db, %s)` does not return rows", query)
	}
}
