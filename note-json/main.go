package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dkr290/go-projects/note-json/note"
)

func main() {

	title, content := getNoteData()
	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	note.PrintNote()
	err = note.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note suceeded")

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
