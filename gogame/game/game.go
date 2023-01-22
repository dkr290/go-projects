package game

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

// errInvlidWorldLength is returned when the guess has wrong number of characters
var errInvlidWorldLength = fmt.Errorf("invalid guess, word does not have the same number of characters as the solution")

type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

func NewGame(playerInput io.Reader, solution string, maxAttempts int) *Game {
	return &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUpperCaseCharacters(solution),
		maxAttempts: maxAttempts,
	}
}

func (g *Game) Play() {
	fmt.Println("Welcome to gessing words game")

	// will ask for valid
	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()
		if slices.Equal(guess, g.solution) {
			fmt.Printf("You won! You found it in %d guess! The word was: %s \n", currentAttempt, string(g.solution))
			return
		}

	}

	fmt.Printf("You have lost! The solution was %s \n", string(g.solution))

}

func (g *Game) ask() []rune {

	fmt.Printf("Please enter %d-character guess:\n", len(g.solution))

	for {
		playerIn, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, "Game failed to read your guess: %w \n", err)
			continue
		}
		guess := splitToUpperCaseCharacters(string(playerIn))

		//verify if the guess is with valid length
		if err := g.validateGuess(guess); err != nil {

			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with the solution %s \n", err.Error())
			return nil
		} else {
			return guess
		}
	}
}

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d characters , got %d, %w", len(g.solution), len(guess), errInvlidWorldLength)
	}

	return nil
}
func splitToUpperCaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}
