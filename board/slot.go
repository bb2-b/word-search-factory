package board

import (
	"fmt"
)

type Slot struct {
	row    int
	col    int
	filled bool
	char   byte
}

func NewSlot(row, col int) *Slot {
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
		s = append(s, *NewSlot(row, col))
	}
	return s
}

func (g *gameBoard) GetRandomSlot() *Slot {
	rowLen := len(g.grid)
	colLen := len(g.grid[0])

	randRow := RandomNumInRange(rowLen)
	randCol := RandomNumInRange(colLen)

	if g.grid[randRow][randCol].filled {
		fmt.Printf("slot was already filled! [%d,%d]\n", randRow, randCol)
		g.GetRandomSlot()
	}

	return &g.grid[randRow][randCol]
}
