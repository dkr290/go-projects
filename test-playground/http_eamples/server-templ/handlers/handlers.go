package handlers

import (
	"log/slog"
	"net/http"
	"server-templ/internal/models"
	"server-templ/templates"
	"strconv"
	"strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) error {
	vl := r.URL.Query()
	cust := models.Customer{}
	var err error
	id, ok := vl["id"]
	if ok {
		cust.Id, err = strconv.Atoi(strings.Join(id, ","))
		if err != nil {
			return err
		}
	}

	name, ok := vl["name"]
	if ok {
		cust.Name = strings.Join(name, ",")
	}

	surname, ok := vl["surname"]
	if ok {
		cust.Surname = strings.Join(surname, ",")
	}

	age, ok := vl["age"]
	if ok {
		cust.Age, err = strconv.Atoi(strings.Join(age, ""))
		if err != nil {
			return err
		}
	}

	return templates.TemplateIndex(cust).Render(r.Context(), w)
}

func MakeHandler(next func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			slog.Error("internal server error", "err", err, "path", r.URL.Path)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
