package main

import (
	"os"

	"github.com/dkr290/go-projects/gogame/game"
)

func main() {

	g := game.NewGame(os.Stdin)
	g.Play()
}
