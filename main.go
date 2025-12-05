package main

import (
	board "github.com/bb2-b/word-search-factory/board"
)

func main() {

	words := []string{
		"random",
		"list",
		"of",
		"words",
	}

	game, err := board.NewGameBoard(words)
}
