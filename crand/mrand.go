package crand

import "math/rand"

// MrandRange gives rand int number in range [min, max]
func MrandRange(min, max int) int {
	return rand.Intn(max-min) + min
}
