package board

func (g *gameBoard) PlaceWords() error {
	for _, vector := range g.wordVectors {
		currSlot := vector.anchor
		for _, char := range vector.word {
			err := g.placeChar(byte(char), currSlot)
			if err != nil {
				return err
			}
			currSlot = updateDirection(vector.direction, currSlot)
		}
	}

	return nil
}
