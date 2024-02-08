package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (t *Todo) Display() {
	fmt.Printf(t.Text)
}

// just save a copy no need pointer
func (t Todo) Save() error {

	fileName := "todo.json"

	jsonContent, err := json.Marshal(&t)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonContent, 0644)

}

func New(c string) (*Todo, error) {

	if c == "" {

		return &Todo{}, errors.New("invalid input")

	}
	return &Todo{
		Text: c,
	}, nil
}
