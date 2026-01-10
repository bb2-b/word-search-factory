package board

import (
	"cmp"
	"fmt"
)

type WordVector struct {
	word      string
	direction direction
	anchor    Slot
}

func (g *gameBoard) newWordVector(word string) *WordVector {
	return &WordVector{
		word:      word,
		anchor:    *g.GetRandomSlot(),
		direction: g.randVector(),
	}
}

func (g *gameBoard) PickWordVector(w string) *WordVector {
	vector := g.newWordVector(w)

	// Ensure that direction can fit the word, recurse until the word can fit.
	err := g.ensureBoardFitness(vector)
	if err != nil {
		return g.PickWordVector(w)
	}

	// Ensure spelling the word into its slots will not have a collision issue.
	err = g.ensureCollisionFitness(vector)
	if err != nil {
		if _, ok := err.(*FilledError); ok {
			return g.PickWordVector(w)
		} else {
			panic("collision problem that was not a FilledError")
		}
	}

	vector.printAnchor()

	return vector
}

func (g *gameBoard) ensureBoardFitness(vector *WordVector) error {
	if vector == nil {
		return fmt.Errorf("provided a nil word vector")
	}

	horErr := g.ensureFitHorizontal(*vector)
	vertErr := g.ensureFitVertical(*vector)

	if err := cmp.Or(horErr, vertErr); err != nil {
		return fmt.Errorf("horizontal(%v), vertical(%v)", horErr, vertErr)
	}

	return nil
}

func (g *gameBoard) ensureFitHorizontal(v WordVector) error {
	switch v.direction {
	case right, upRight, downRight:
		if v.anchor.col+len(v.word)-1 > len(g.grid[0])-1 {
			return fmt.Errorf("word went right off the board")
		}
	case left, upLeft, downLeft:
		if v.anchor.col-len(v.word)-1 < 0 {
			return fmt.Errorf("word went left off the board")
		}
	}

	return nil
}

func (g *gameBoard) ensureFitVertical(v WordVector) error {
	switch v.direction {
	case down, downLeft, downRight:
		if v.anchor.row+len(v.word)-1 > len(g.grid[0])-1 {
			return fmt.Errorf("word went down off the board")
		}
	case up, upLeft, upRight:
		if v.anchor.row-len(v.word)-1 < 0 {
			return fmt.Errorf("word went up off the board")
		}
	}

	return nil
}

func (g *gameBoard) ensureCollisionFitness(v *WordVector) error {
	currSlot := v.anchor
	for _, char := range v.word {
		err := g.fauxPlaceChar(byte(char), currSlot)
		if err != nil {
			return err
		}
		currSlot = updateDirection(v.direction, currSlot)
	}

	return nil
}

func updateDirection(d direction, s Slot) Slot {
	switch d {
	case upLeft:
		s.row -= 1
		s.col -= 1
	case up:
		s.row -= 1
	case upRight:
		s.row -= 1
		s.col += 1
	case left:
		s.col -= 1
	case right:
		s.col += 1
	case downLeft:
		s.row += 1
		s.col -= 1
	case down:
		s.row += 1
	case downRight:
		s.col += 1
		s.row += 1
	}

	return s
}

func (g *gameBoard) fauxPlaceChar(char byte, slot Slot) *FilledError {
	gSlot := &g.grid[slot.row][slot.col]
	if !gSlot.filled {
		return nil
	}

	// If the designated slot happened to contain the same letter already...
	if gSlot.char == char {
		return nil
	}

	return NewFilledError(slot)
}

func (v *WordVector) printAnchor() {
	fmt.Printf("anchor [r%.2d,c%.2d] (%s) - '%s'\n",
		v.anchor.row, v.anchor.col, v.direction, v.word)
}
