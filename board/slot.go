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

	rRow := RandomNumInRange(rowLen)
	rCol := RandomNumInRange(colLen)

	if g.grid[rRow][rCol].filled {
		fmt.Printf("That slot was already filled! [%d,%d]\n", rRow, rCol)
		g.GetRandomSlot()
	}

	return &g.grid[rRow][rCol]
}
