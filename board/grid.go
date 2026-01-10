package board

import (
	"fmt"
	"math/rand"
	"strings"
)

// type GameBoard interface {
// 	// NewGameBoard returns a new instance of a game board or an error.
// 	NewGameBoard(words []string) (*gameBoard, error)

// 	// GetGrid returns the grid of a game board.
// 	Grid() [][]Slot

// 	// GetWordList returns the word list composing a game board.
// 	WordList() []string

// 	// PrettyPrintGameBoard prints a pretty game board.
// 	PrettyPrintGameBoard()
// }

type gameBoard struct {
	grid        [][]Slot
	difficulty  difficulty
	wordList    []string
	wordVectors []*WordVector
	key         [][]Slot
}

func (g *gameBoard) Grid() [][]Slot {
	return g.grid
}

func (g *gameBoard) WordList() []string {
	return g.wordList
}

func (g *gameBoard) Key() [][]Slot {
	return g.key
}

func NewGameBoard(words *[]string, difficulty string) (*gameBoard, error) {
	if words == nil {
		return nil, fmt.Errorf("provided null word list")
	}

	// Create empty game.
	size, err := getMinBoardSize(*words)
	if err != nil {
		return nil, err
	}
	game := &gameBoard{
		grid:       createGrid(size),
		difficulty: DifficultyStrToInt(difficulty),
		wordList:   *words,
	}

	// Generate vector for each word and place words into its position.
	for _, word := range game.wordList {
		word = strings.ToUpper(word)

		vec := game.PickWordVector(word)
		game.wordVectors = append(game.wordVectors, vec)

		// Place words into their slots.
		err = game.PlaceWords()
		if err != nil {
			return game, err
		}
	}

	// Save the game to serve as the answer key.
	game.deepCopyGrid()

	// Fill the unfilled slots.
	game.randomlyFill()

	return game, nil
}

// getMinBoardSize returns the minimum size of the game Board required to fit
// the longest word.
func getMinBoardSize(words []string) (int, error) {
	if len(words) == 0 {
		return 0, fmt.Errorf("no words provided in word list")
	}

	var maxWordLen int
	for _, word := range words {
		if len(word) > maxWordLen {
			// Ensure the board can fit the longest word plus the number of
			// words in the word list.
			maxWordLen = len(word) + len(words) - 1
		}
	}

	return maxWordLen, nil
}

// createGrid returns an (n x n) 2d slice of bytes (underscores).
// The size of the board is based on the longest word in the provided list plus
// the number of remaining words.
func createGrid(length int) [][]Slot {
	// TODO: choosing board size (see above comment)
	rows, cols := length, length
	grid := make([][]Slot, rows)
	for row := range grid {
		grid[row] = NewSetOfSlots(row, cols)
	}

	fmt.Printf("created %dx%d grid\n", length, length)

	return grid
}

func (g *gameBoard) deepCopyGrid() {
	g.key = make([][]Slot, len(g.grid))
	for i := range g.grid {
		g.key[i] = make([]Slot, len(g.grid[i]))
		copy(g.key[i], g.grid[i])
	}
}

func (g *gameBoard) randomlyFill() {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, col := range g.grid {
		for _, slot := range col {
			if !slot.filled {
				rChar := charset[rand.Intn(len(charset))]
				g.placeChar(byte(rChar), slot)
			}
		}
	}
}

func (g *gameBoard) PrettyPrintGameBoard(board [][]Slot) {
	fmt.Printf("difficulty: %s\n", g.difficulty.String())

	fmt.Print("  ")
	for i := range board[0] {
		fmt.Printf("%-3d", i)
	}
	fmt.Println()
	for j, row := range board {
		fmt.Printf("%-2d", j)
		for _, slot := range row {
			fmt.Printf("%-*s", 3, string(slot.char))
		}
		fmt.Println()
	}
}
