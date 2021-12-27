package rand

import (
	"math"
	"math/rand"
)

func RandomWithLength(length int) int {
	min := int(math.Pow10(length - 1))
	max := int(math.Pow10(length)) - 1
	return rand.Intn(max-min) + min
}
