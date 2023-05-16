package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Helpers struct{}

func New() *Helpers {
	return &Helpers{}
}

func (h *Helpers) ReadJsonFromHttp(w http.ResponseWriter, r *http.Request, data any) error {

	maxBytes := 1048576 //one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	log.Println(r.Body)

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&data); err != nil {
		return err

	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must have only a single JSON value")

	}

	return nil
}

func (h *Helpers) SendJsonResponse(w http.ResponseWriter, status int, data any) error {

	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil

}
