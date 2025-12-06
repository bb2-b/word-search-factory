package board

import (
	"fmt"
)

type GameBoard interface {
	// NewGameBoard returns a new instance of a game board or an error.
	NewGameBoard(words []string) (*gameBoard, error)

	// GetGrid returns the grid of a game board.
	Grid() [][]Slot

	// GetWordList returns the word list composing a game board.
	WordList() []string

	// PrettyPrintGameBoard prints a pretty game board.
	PrettyPrintGameBoard()
}

type gameBoard struct {
	grid     [][]Slot
	wordList []string
}

func (g *gameBoard) Grid() [][]Slot {
	return g.grid
}

func (g *gameBoard) WordList() []string {
	return g.wordList
}

func NewGameBoard(words *[]string) (*gameBoard, error) {
	if words == nil {
		return nil, fmt.Errorf("provided nil word list")
	}

	// Create empty game.
	size, err := getMinBoardSize(*words)
	if err != nil {
		return nil, err
	}

	grid := createGrid(size)

	game := &gameBoard{
		grid:     grid,
		wordList: *words,
	}

	// Randomly pick an anchor point for each word.
	for _, word := range game.wordList {
		game.PickWordAnchor(word)
	}

	return game, nil
}

func (g *gameBoard) PrettyPrintGameBoard() {
	for _, row := range g.grid {
		for _, slot := range row {
			fmt.Printf("%-*s", 2, string(slot.char))
		}
		fmt.Println()
	}
}

// getMinBoardSize returns the minimum size of the game Board required to fit the longest word.
func getMinBoardSize(words []string) (int, error) {
	if len(words) == 0 {
		return 0, fmt.Errorf("no words provided in word list")
	}

	var maxWordLen int
	for _, word := range words {
		if len(word) > maxWordLen {
			maxWordLen = len(word)
		}
	}

	return maxWordLen, nil
}

// createGrid returns an (n x n) 2d slice of bytes (underscores).
func createGrid(length int) [][]Slot {
	rows, cols := length, length
	grid := make([][]Slot, rows)
	for row := range grid {
		grid[row] = NewSetOfSlots(row, cols)
	}

	return grid
}
