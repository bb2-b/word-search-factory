package board

import (
	"cmp"
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
	err := g.ensureFit(vector)
	if err != nil {
		fmt.Printf("fitment problem: %s\n", err.Error())
		vector = g.PickWordVector(w)
	}

	// Ensure spelling the word into its slots will not have a collision issue.
	err = g.PreCheckFit(vector)

	return vector
}

func (g *gameBoard) PlaceWords() error {
	fmt.Printf("word vectors: %+v\n", g.wordVectors)
	for _, vector := range g.wordVectors {

		currChar := vector.anchor
		for _, char := range vector.word {
			err := g.slotInPlace(byte(char), currChar)
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

func (g *gameBoard) slotInPlace(char byte, slot Slot) error {
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

// TODO: direction generation based on difficulty
func (g *gameBoard) randVector() direction {
	return direction(rand.Intn(8))
}

func (g *gameBoard) ensureFit(vector *WordVector) error {
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
	fmt.Printf("anchor [r%.2d,c%.2d] (%s) - '%s'\n",
		v.anchor.row, v.anchor.col, v.direction, v.word)
}

func (g *gameBoard) PreCheckFit(vector *WordVector) error {
	currChar := vector.anchor
	for _, char := range vector.word {
		err := g.fauxSlotInPlace(byte(char), currChar)
		if err != nil {
			if _, ok := err.(*FilledError); ok {
				panic("filled error!")
				// g.PickWordVector(vector.word)
				// break
			}
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

	return nil
}

func (g *gameBoard) fauxSlotInPlace(char byte, slot Slot) error {
	gSlot := &g.grid[slot.row][slot.col]
	if !gSlot.filled {
		return nil
	} else {
		// If the designated slot happened to contain the same letter already...
		if gSlot.char == char {
			fmt.Printf("coincidental matching letter!!!\n\n")
			return nil
		} else {
			return NewFilledError()
		}
	}

	fmt.Printf("fake placed letter\n")
	fmt.Printf("\n\n")
	return nil
}
