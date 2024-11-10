package handlers

import (
	"fmt"
	"net/http"
	"prod/db"
)

func HandleIndex(w http.ResponseWriter, r *http.Request, db db.Database) error {
	fmt.Fprint(w, "This is a test index function")
	if err := db.GetDB(); err != nil {
		return err
	}
	return nil
}
