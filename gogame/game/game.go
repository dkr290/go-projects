package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const wordLength = 5

type Game struct {
	reader *bufio.Reader
}

func NewGame(playerInput io.Reader) *Game {
	return &Game{
		reader: bufio.NewReader(playerInput),
	}
}

func (g *Game) Play() {
	fmt.Println("Welcome to gessing words game")

	// will ask for valid word
	guess := g.ask()

	fmt.Println("Your guess is ", string(guess))

}

func (g *Game) ask() []rune {

	fmt.Printf("Please enter %d-character guess:\n", wordLength)

	for {
		playerIn, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, "Game failed to read your guess: %w \n", err)
			continue
		}
		guess := []rune(string(playerIn))

		//verify if the guess is with valid length
		if len(guess) != wordLength {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid. Expected %d characters , got %d", wordLength, len(guess))
		} else {
			return guess
		}
	}
}
