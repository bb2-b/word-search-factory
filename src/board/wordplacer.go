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
	vector.printAnchor()

	// Ensure that direction can fit the word, recurse until the word can fit.
	err := g.ensureBoardFitness(vector)
	if err != nil {
		fmt.Printf("fitment problem: %s\n", err.Error())
		vector = g.PickWordVector(w)
	}

	// Ensure spelling the word into its slots will not have a collision issue.
	err = g.ensureCollisionFitness(vector)
	if err != nil {
		fmt.Printf("fitment problem: %s\n", err.Error())
		vector = g.PickWordVector(w)
	}

	return vector
}

func (g *gameBoard) PlaceWords() error {
	fmt.Printf("word vectors: %+v\n", g.wordVectors)
	for _, vector := range g.wordVectors {
		currChar := vector.anchor
		for _, char := range vector.word {
			err := g.placeChar(byte(char), currChar)
			if err != nil {
				return err
			}
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

	return nil
}

// func (g *gameBoard) fauxPlaceChar(char byte, slot Slot) error {
// 	gSlot := &g.grid[slot.row][slot.col]
// 	if !gSlot.filled {
// 		return nil
// 	}

// 	// If the designated slot happened to contain the same letter already...
// 	if gSlot.char == char {
// 		fmt.Printf("faux: coincidental matching letter!!!\n\n")
// 		return nil
// 	}

// 	return NewFilledError()
// }

func (v *WordVector) printAnchor() {
	fmt.Printf("anchor [r%.2d,c%.2d] (%s) - '%s'\n",
		v.anchor.row, v.anchor.col, v.direction, v.word)
}
