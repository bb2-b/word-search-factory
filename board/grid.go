package board

import (
	"fmt"
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
	wordList    []string
	wordVectors []*WordVector
}

func (g *gameBoard) Grid() [][]Slot {
	return g.grid
}

func (g *gameBoard) WordList() []string {
	return g.wordList
}

func NewGameBoard(words *[]string) (*gameBoard, error) {
	if words == nil {
		return nil, fmt.Errorf("provided null word list")
	}

	// Create empty game.
	size, err := getMinBoardSize(*words)
	if err != nil {
		return nil, err
	}

	grid := createGrid(size)

	game := &gameBoard{
		grid:        grid,
		wordList:    *words,
		wordVectors: nil,
	}

	// Generate vector for each word.
	for _, word := range game.wordList {
		vec := game.PickWordVector(word)
		fmt.Printf("adding word vector to board: %#v\n", vec)
		game.wordVectors = append(game.wordVectors, vec)
	}

	// Place words into their slots.
	err = game.PlaceWords()
	if err != nil {
		return game, err
	}

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

func (g *gameBoard) PlaceWords() error {
	fmt.Printf("word vectors: %s\n", g.wordVectors)
	for _, vector := range g.wordVectors {

		currChar := vector.anchor
		for _, char := range vector.word {
			err := g.slotInPlace(byte(char), currChar)
			if err != nil {
				return err
			}
			fmt.Printf("vector: %#v\n", vector)
			switch vector.direction {
			case upLeft:
				currChar.row -= 1
				currChar.col -= 1
			case up:
				currChar.row -= 1
			case upRight:
				currChar.row -= 1
				currChar.col += 1
			case left:
				currChar.col -= 1
			case right:
				currChar.col += 1
			case downLeft:
				currChar.row += 1
				currChar.col -= 1
			case down:
				currChar.row += 1
			case downRight:
				currChar.col += 1
				currChar.row += 1
			default:
				return fmt.Errorf("something went wrong adjusting the follow-on character direction")
			}
		}
	}

	return nil
}

func (g *gameBoard) slotInPlace(char byte, slot Slot) error {
	gSlot := &g.grid[slot.row][slot.col]
	if !gSlot.filled {
		gSlot.char = char
		gSlot.filled = true
	} else {
		return fmt.Errorf("slot was already filled")
	}

	g.PrettyPrintGameBoard()
	fmt.Printf("\n\n")
	return nil
}

func (g *gameBoard) PrettyPrintGameBoard() {
	fmt.Print("  ")
	for i := range g.grid[0] {
		fmt.Printf("%-2d", i)
	}
	fmt.Println()
	for j, row := range g.grid {
		fmt.Printf("%-2d", j)
		for _, slot := range row {
			fmt.Printf("%-*s", 2, string(slot.char))
		}
		fmt.Println()
	}
}
