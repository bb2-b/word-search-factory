package board

import "math/rand"

func RandomNumInRange(n int) int {
	return rand.Intn(n)
}
