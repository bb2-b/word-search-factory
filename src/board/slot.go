package board

import (
	"fmt"
	"math/rand"
)

type Slot struct {
	row    int
	col    int
	filled bool
	char   byte
}

type FilledError struct{}

func newSlot(row, col int) *Slot {
	s := &Slot{
		filled: false,
		row:    row,
		col:    col,
		char:   '_',
	}
	return s
}

func NewSetOfSlots(row, numOfCols int) (s []Slot) {
	for col := range numOfCols {
		s = append(s, *newSlot(row, col))
	}
	return s
}

func (g *gameBoard) GetRandomSlot() *Slot {
	rowLen := len(g.grid)
	colLen := len(g.grid[0])

	randRow := rand.Intn(rowLen)
	randCol := rand.Intn(colLen)

	if g.grid[randRow][randCol].filled {
		fmt.Printf("slot was already filled! [%d,%d]\n", randRow, randCol)
		g.GetRandomSlot()
	}

	return &g.grid[randRow][randCol]
}

func (g *gameBoard) placeChar(char byte, slot Slot) error {
	gSlot := &g.grid[slot.row][slot.col]
	if !gSlot.filled {
		gSlot.char = char
		gSlot.filled = true
	} else {
		// If the designated slot happened to contain the same letter already...
		if gSlot.char == char {
			fmt.Printf("coincidental matching letter!!!\n\n")
			return nil
		}
		return fmt.Errorf("slot was already filled with unmatched letter")
	}

	g.PrettyPrintGameBoard()
	fmt.Printf("\n\n")
	return nil
}

// Error implements the error interface.
func (e *FilledError) Error() string {
	return fmt.Sprintf("resource '%s' is filled", e.Error())
}

// NewFilledError creates a new FilledError
func NewFilledError() *FilledError {
	return &FilledError{}
}
