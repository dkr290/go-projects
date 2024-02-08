package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *Note) Display() {
	fmt.Printf("The note title %s has the following content:\n\n%s\n", n.Title, n.Content)
}

// just save a copy no need pointer
func (n Note) Save() error {

	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	jsonContent, err := json.Marshal(&n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonContent, 0644)

}

func New(t, c string) (*Note, error) {

	if t == "" || c == "" {

		return &Note{}, errors.New("invalid input")

	}
	return &Note{
		Title:     t,
		Content:   c,
		CreatedAt: time.Now(),
	}, nil
}
