package main

import (
	board "github.com/bb2-b/word-search-factory/board/board"
)

func main() {

	words := []string{
		"random",
		"list",
		"words",
		"lorem",
		"ipsum",
		"brandon",
		"words",
		// "supercalifragilistic",
		// "supercalifragilisticexpialidocious",
	}

	// game, err := board.NewGameBoard(&words, "easy")
	// game, err := board.NewGameBoard(&words, "medium")
	game, err := board.NewGameBoard(&words, "hard")
	// game, err := board.NewGameBoard(&words, "extreme")
	if err != nil {
		panic(err)
	}

	game.PrettyPrintGameBoard(game.Grid())
	game.PrettyPrintGameBoard(game.Key())
}
