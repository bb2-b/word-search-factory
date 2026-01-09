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

// TODO: direction generation based on difficulty
func (g *gameBoard) randVector() direction {
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
		return fmt.Sprintf("Unknown direction (%d)", d)
	}
}
