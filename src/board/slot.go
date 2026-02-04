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

type FilledError struct {
	Slot
}

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

func (g *gameBoard) GetRandomSlot() Slot {
	rowLen := len(g.grid)
	colLen := len(g.grid[0])

	randRow := rand.Intn(rowLen)
	randCol := rand.Intn(colLen)

	if g.grid[randRow][randCol].filled {
		g.GetRandomSlot()
	}

	return g.grid[randRow][randCol]
}

func (g *gameBoard) placeChar(char byte, slot Slot) error {
	gSlot := &g.grid[slot.row][slot.col]
	if !gSlot.filled {
		gSlot.char = char
		gSlot.filled = true
	} else {
		// If the designated slot happened to contain the same letter already...
		if gSlot.char == char {
			return nil
		}
		return fmt.Errorf("slot was already filled with unmatched letter")
	}

	return nil
}

// Error implements the error interface.
func (e *FilledError) Error() string {
	return fmt.Sprintf("resource '%+v' is filled", e.Slot)
}

// NewFilledError creates a new FilledError
func NewFilledError(s Slot) *FilledError {
	return &FilledError{Slot: s}
}
