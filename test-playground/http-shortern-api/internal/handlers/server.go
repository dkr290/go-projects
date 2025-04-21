package handlers

import (
	"errors"
	"fmt"
	"http-shortern-api/internal/store"
	"http-shortern-api/internal/utils"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer(s *store.Store) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /shorten", Shorten(s))
	mux.HandleFunc("GET /r/{key}", Resolve(s))
	mux.HandleFunc("GET /health", Health)

	return &Server{
		mux: mux,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func Shorten(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		link := store.Link{
			Key: r.FormValue("key"),
			URL: r.FormValue("url"),
		}
		if err := s.Create(r.Context(), link); err != nil {
			httpError(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(link.Key))
	}
}

func Resolve(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		link, err := s.Retrieve(r.Context(), r.PathValue("key"))
		if err != nil {
			httpError(w, err)
			return
		}

		http.Redirect(w, r, link.URL, http.StatusFound)
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func httpError(w http.ResponseWriter, err error) {
	if err == nil { // no error
		return
	}
	var code int
	switch {
	case errors.Is(err, utils.ErrInvalidRequest):
		code = http.StatusBadRequest
	case errors.Is(err, utils.ErrExists):
		code = http.StatusConflict
	case errors.Is(err, utils.ErrNotExist):
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}
	http.Error(w, err.Error(), code)
}
