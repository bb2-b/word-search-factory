package board

import (
	"fmt"
	"math/rand"
	"strings"
)

type gameBoard struct {
	difficulty
	grid        [][]Slot
	wordList    []string
	wordVectors []*WordVector
	key         [][]Slot
}

// Grid returns a 2d slice of slots that contains the instance of a game board
// at any given time.
func (g *gameBoard) Grid() [][]Slot {
	return g.grid
}

// WordList returns a slice of strings containing the list of words the game
// board will be built from.
func (g *gameBoard) WordList() []string {
	return g.wordList
}

// AnswerKey returns a 2d slice of (un-filled) slots with the placed words from
// the game's word list.
func (g *gameBoard) AnswerKey() [][]Slot {
	return g.key
}

func NewGameBoard(words []string, difficulty string) (*gameBoard, error) {
	if words == nil {
		return nil, fmt.Errorf("provided null word list")
	}

	// Create empty game.
	size, err := getMinBoardSize(words)
	if err != nil {
		return nil, err
	}
	game := &gameBoard{
		grid:       createGrid(size),
		difficulty: DifficultyStrToInt(difficulty),
		wordList:   words,
	}

	// Generate vector for each word and place words into its position.
	for _, word := range game.wordList {
		word = strings.ToUpper(word)
		word = strings.TrimSpace(word)

		game.wordVectors = append(game.wordVectors, game.PickWordVector(word))

		// Place words into their designated slots.
		err = game.PlaceWords()
		if err != nil {
			return game, err
		}
	}

	// Save the game to serve as the answer key.
	game.deepCopyGrid()

	// Fill the unfilled slots.
	game.randFill()

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

// createGrid returns an (n x n) 2d slice of bytes comprised of underscores.
// The size of the board is based on the longest word in the provided list plus
// the number of remaining words, subsequently, ensuring all words can fit on
// the board.
func createGrid(length int) [][]Slot {
	rows, cols := length, length
	grid := make([][]Slot, rows)
	for row := range grid {
		grid[row] = NewSetOfSlots(row, cols)
	}

	return grid
}

func (g *gameBoard) PlaceWords() error {
	for _, vector := range g.wordVectors {
		currSlot := vector.Slot
		for _, char := range vector.word {
			err := g.placeChar(byte(char), currSlot)
			if err != nil {
				return err
			}
			currSlot = updateDirection(vector.direction, currSlot)
		}
	}

	return nil
}

func (g *gameBoard) deepCopyGrid() {
	g.key = make([][]Slot, len(g.grid))
	for i := range g.grid {
		g.key[i] = make([]Slot, len(g.grid[i]))
		copy(g.key[i], g.grid[i])
	}
}

func (g *gameBoard) randFill() {
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

func (g *gameBoard) PrettyPrintGameBoard(board [][]Slot, showCoord bool) {
	fmt.Printf("difficulty: %s\n", g.difficulty.String())

	if showCoord { // corresponds to the board's columns
		fmt.Print("  ")
		for i := range board[0] {
			fmt.Printf("%-3d", i)
		}
		fmt.Println()
	}
	for j, row := range board {
		if showCoord { // corresponds to the board's rows
			fmt.Printf("%-2d", j)
		}
		for _, slot := range row {
			fmt.Printf("%-*s", 3, string(slot.char))
		}
		fmt.Println()
	}
}
