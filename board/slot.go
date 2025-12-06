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

func randomNumInRange(n int) int {
	return rand.Intn(n)
}

func (g *gameBoard) GetRandomSlot() *Slot {
	rowLen := len(g.grid)
	colLen := len(g.grid[0])

	rRow := randomNumInRange(rowLen)
	rCol := randomNumInRange(colLen)

	if g.grid[rRow][rCol].filled {
		fmt.Printf("No way...that slot was already filled! [%d,%d]\n", rRow, rCol)
		g.GetRandomSlot()
	}

	return &g.grid[rRow][rCol]
}
