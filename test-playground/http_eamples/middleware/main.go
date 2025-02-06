package main

import (
	"log"
	"log/slog"
	"net/http"
)

func Middle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "this is from middleware"
		_, err := w.Write([]byte(msg))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return

		}
		next.ServeHTTP(w, r)
	}
}

func MakeHandler(next func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			slog.Error("internal server error", "err", err, "path", r.URL.Path)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Func1(w http.ResponseWriter, r *http.Request) error {
	msg := "This is from the first function"
	if _, err := w.Write([]byte(msg)); err != nil {
		return err
	}
	return nil
}

func Func2(w http.ResponseWriter, r *http.Request) {
	msg := "This is from the second function"
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/f1", MakeHandler(Func1))
	http.HandleFunc("/f2", Middle(Func2))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
