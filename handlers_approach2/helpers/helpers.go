package helpers

import (
	"log/slog"
	"net/http"
	"prod/db"
)

func MakeHandlers(
	handlerFunc func(w http.ResponseWriter, r *http.Request, db db.Database) error,
	db db.Database,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handlerFunc(w, r, db); err != nil {
			slog.Error("Internal server error", "err", err, "path", r.URL.Path)
		}
	}
}
