package board

import (
	"fmt"
	"math/rand"
)

const (
	right     direction = iota // 0 - easy
	down                       // 1 - easy
	up                         // 2 - medium
	upRight                    // 3 - medium
	downRight                  // 4 - medium
	upLeft                     // 5 - hard
	left                       // 6 - hard
	downLeft                   // 7 - hard
)

type direction int

func (g *gameBoard) randVector() direction {
	switch g.difficulty {
	case easy:
		return direction(rand.Intn(2))
	case medium:
		return direction(rand.Intn(4))
	case hard:
		return direction(rand.Intn(8))
	}

	fmt.Println("defaulting to 'hard' difficulty")
	return direction(rand.Intn(8))
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
		return fmt.Sprintf("unknown direction (%d)", d)
	}
}
