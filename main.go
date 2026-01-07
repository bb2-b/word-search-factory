package main

import (
	board "github.com/bb2-b/word-search-factory/board/board"
)

func main() {

	words := []string{
		"random",
		"random",
		// "list",
		// "of",
		// "words",
		// "supercalifragilistic",
		// "supercalifragilisticexpialidocious",
	}

	game, err := board.NewGameBoard(&words)
	if err != nil {
		panic(err)
	}

	game.PrettyPrintGameBoard()
}
