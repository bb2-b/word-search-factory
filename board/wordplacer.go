package board

import (
	"fmt"
)

func (g *gameBoard) PlaceWords() error {
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
