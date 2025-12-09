package board

import "fmt"

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
	direction direction
	slot      Slot
	word      string
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

func (g *gameBoard) EnsureFit(a WordVector) error {
	switch a.direction {
	// Check horizontally (col).
	case upLeft, left, downLeft:
		if a.slot.col-len(a.word)-1 < 0 {
			return fmt.Errorf("word went left off the board")
		}
	case upRight, right, downRight:
		if a.slot.col+len(a.word)-1 > len(g.grid[0]) {
			return fmt.Errorf("word went right off the board")
		}
	// Check vertically (row).
	case up:
		if a.slot.row-len(a.word)-1 < 0 {
			return fmt.Errorf("word went up off the board")
		}
	case down:
		if a.slot.row+len(a.word)-1 > len(g.grid[0]) {
			return fmt.Errorf("word went down off the board")
		}
	}

	return nil
}

func (g *gameBoard) RandVector(a WordVector) direction {
	return direction(RandomNumInRange(8))
}

func (g *gameBoard) NewWordVector(word string) *WordVector {
	s := g.GetRandomSlot()
	return &WordVector{
		slot: *s,
		word: word,
	}
}

func (g *gameBoard) PickWordAnchor(w string) *WordVector {
	// Pick first letter anchor.
	anchor := g.NewWordVector(w)
	anchor.printAnchor()

	// Choose random direction for word to be spelled into.
	anchor.direction = g.RandVector(*anchor)
	fmt.Printf("direction: %v\n", anchor.direction.String())

	// Ensure that direction can fit the word.
	err := g.EnsureFit(*anchor)
	if err != nil {
		fmt.Printf("fitment problem: %s\n", err.Error())
		// Grid has enough slots to ensure infinite recursion cannot happen.
		g.PickWordAnchor(w)
	}

	return anchor
}

func (a *WordVector) printAnchor() {
	fmt.Printf("[%.2d,%.2d] - '%s'\n", a.slot.row, a.slot.col, a.word)
}
