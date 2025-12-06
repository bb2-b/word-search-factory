package board

import "fmt"

type Anchor struct {
	slot Slot
	word string
}

type direction int

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

// String method for direction
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

// TODO
func (g *gameBoard) RandWordVector() direction {
	return 0
}

func (g *gameBoard) NewAnchor(word string) *Anchor {
	s := g.GetRandomSlot()
	return &Anchor{
		slot: *s,
		word: word,
	}
}

func (g *gameBoard) PickWordAnchor(word string) *Anchor {
	// Pick first letter anchor.
	a := g.NewAnchor(word)
	a.printAnchor()

	// Choose random direction for word to be spelled.
	// TODO: direction := g.RandWordVector()

	// Ensure that vector can fit the word.

	return nil
}

func (a *Anchor) printAnchor() {
	fmt.Printf("[%.2d,%.2d] - '%s'\n", a.slot.row, a.slot.col, a.word)
}
