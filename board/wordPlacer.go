package board

import (
	"fmt"
	"math/rand"
)

const (
	upLeft    direction = iota // 0
	up                         // 1
	upRight                    // 2
	left                       // 3
	right                      // 4
	downLeft                   // 5
	down                       // 6
	downRight                  // 7
)

type direction int

type WordVector struct {
	word      string
	direction direction
	anchor    Slot
}

func (g *gameBoard) PickWordVector(w string) *WordVector {
	vector := g.newWordVector(w)
	vector.printAnchor()

	fmt.Printf("direction: %v\n", vector.direction.String())

	// Ensure that direction can fit the word, recurse until the word can fit.
	err := g.ensureFit(*vector)
	if err != nil {
		fmt.Printf("fitment problem: %s\n", err.Error())
		g.PickWordVector(w)
	}

	fmt.Printf("returning new vector: %#v\n", vector)
	return vector
}

func (g *gameBoard) newWordVector(word string) *WordVector {
	return &WordVector{
		word:      word,
		anchor:    *g.GetRandomSlot(),
		direction: g.randVector(),
	}
}

// TODO: direction generation based on difficulty
func (g *gameBoard) randVector() direction {
	return direction(rand.Intn(8))
}

func (g *gameBoard) ensureFit(anch WordVector) error {
	horErr := g.ensureFitHorizontal(anch)
	vertErr := g.ensureFitVertical(anch)

	if horErr != nil || vertErr != nil {
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

// String method for direction.
func (d direction) String() string {
	switch d {
	case upLeft:
		return "up left"
	case up:
		return "up"
	case upRight:
		return "up right"
	case left:
		return "left"
	case right:
		return "right"
	case downLeft:
		return "down left"
	case down:
		return "down"
	case downRight:
		return "down right"
	default:
		return fmt.Sprintf("Unknown direction (%d)", d)
	}
}

func (v *WordVector) printAnchor() {
	fmt.Printf("[r%.2d,c%.2d] - '%s'\n", v.anchor.row, v.anchor.col, v.word)
}
