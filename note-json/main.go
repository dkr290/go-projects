package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dkr290/go-projects/note-json/note"
	"github.com/dkr290/go-projects/note-json/todo"
)

type Saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// interface embedding
type outputtable interface {
	Saver
	Display()
}

func main() {

	printSomething(1)
	printSomething(1.8)
	printSomething("some text")

	title, content := getNoteData()

	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}
	printSomething(todo)

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(note)
	if err != nil {
		return
	}

}

func printSomething(value any) {

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// default:
	// 	// for other types
	// }

	intValue, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intValue)
		return
	}

	floatValue, ok := value.(float64)
	if ok {
		fmt.Println("Integer:", floatValue)
		return
	}

	strngValue, ok := value.(string)
	if ok {
		fmt.Println(strngValue)
		return
	}

}

func outputData(data outputtable) error {

	data.Display()
	return saveData(data)

}

// using the interface
func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the todo failed")
		return err
	}
	fmt.Println("Saving the todo suceeded")
	return nil

}

func getNoteData() (string, string) {

	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")

	return title, content

}
func getUserInput(prompt string) string {
	fmt.Printf("%s ", prompt)

	r := bufio.NewReader(os.Stdin)
	text, err := r.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")

	text = strings.TrimSuffix(text, "\r")

	return text
}
