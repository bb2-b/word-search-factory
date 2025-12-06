package board

import "fmt"

type Anchor struct {
	slot Slot
	row  int
	col  int
	word string
}

func (a *Anchor) PrintAnchor() {
	fmt.Printf("%-2d,%-2d - '%s'\n", a.row, a.col, a.word)
}

func (g *gameBoard) NewAnchor(word string) *Anchor {
	s := g.GetRandomSlot()
	return &Anchor{
		slot: *s,
		row:  s.row,
		col:  s.col,
		word: word,
	}
}

func (g *gameBoard) PickWordAnchor(word string) *Anchor {
	a := g.NewAnchor(word)

	a.PrintAnchor()

	return nil
}
