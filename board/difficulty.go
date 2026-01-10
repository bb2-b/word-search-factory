package board

import "fmt"

type difficulty int

const (
	easy   difficulty = iota // 0
	medium                   // 1
	hard                     // 2
)

func DifficultyStrToInt(diff string) difficulty {
	switch diff {
	case "easy":
		return difficulty(0)
	case "medium":
		return difficulty(1)
	case "hard":
		return difficulty(2)
	default:
		fmt.Println("provided unknown difficulty, defaulting to 'hard'")
		return difficulty(2)
	}
}

// String method for difficulty.
func (d difficulty) String() string {
	switch d {
	case easy:
		return "easy (left to right & up to down)"
	case medium:
		return "medium (left to right on the diagonal & up or down)"
	case hard:
		return "hard (any direction)"
	default:
		return fmt.Sprintf("unknown difficulty (%d)", d)
	}
}
