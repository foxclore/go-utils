package crand

import (
	urand "github.com/foxclore/go-utils/rand"
	"math"
	"math/big"
	mrand "math/rand"
)
import "crypto/rand"

func RandomWithLength(length int) int64 {
	min := int64(math.Pow10(length - 1))
	max := int64(math.Pow10(length)) - 1
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		// Fallback to standard rand:
		return int64(urand.RandomWithLength(length))
	}
	return n.Int64() + min
}

func Random(lens ...int) int64 {
	if len(lens) > 0 {
		return RandomWithLength(lens[0])
	} else {
		return RandomWithLength(mrand.Int())
	}
}

func RandomWithLengthNF(length int) (int64, error) {
	min := int64(math.Pow10(length - 1))
	max := int64(math.Pow10(length)) - 1
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	return n.Int64() + min, err
}

func RandomNoFallback(lens ...int) (int64, error) {
	if len(lens) > 0 {
		return RandomWithLengthNF(lens[0])
	} else {
		return RandomWithLengthNF(mrand.Int())
	}
}
