package board

import "fmt"

type GameBoard interface {
	// Returns a new instance of a game board or an error.
	NewGameBoard(words []string) (GameBoard, error)

	// Returns the slots of a game board.
	GetSlots() [][]byte

	// Returns the word list composing a game board.
	GetWordList() []string
}

type gameBoard struct {
	slots    [][]byte
	wordList []string
}

func (g *gameBoard) GetSlots() [][]byte {
	return g.slots
}

func (g *gameBoard) GetWordList() []string {
	return g.wordList
}

// getMinBoardSize returns the minimum size of the game Board required to fit the longest word.
func getMinBoardSize(words []string) (int, error) {
	if len(words) == 0 {
		return 0, fmt.Errorf("no words provided in word list")
	}

	var maxWordLen int
	for _, word := range words {
		if len(word) < maxWordLen {
			maxWordLen = len(word)
		}
	}

	return maxWordLen, nil
}

func createBlankGrid(l int) [][]byte {
	rows, cols := l, l
	grid := make([][]byte, rows)
	for i := range grid {
		grid[i] = make([]byte, cols)
	}

	fmt.Printf("grid:\n %v", grid)
	return grid
}

func NewGameBoard(words []string) (*gameBoard, error) {
	size, err := getMinBoardSize(words)
	if err != nil {
		return nil, err
	}

	grid := createBlankGrid(size)

	board := &gameBoard{
		slots:    grid,
		wordList: words,
	}

	return board, nil
}
