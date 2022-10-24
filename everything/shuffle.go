package everything

import (
	"github.com/foxclore/go-utils/crand"
)

// simple linear shuffle
func newMatcher(length int) func(int) int {
	a := crand.MrandRange(1, 1000)
	// Here I got a little lesson in trickery: we need to check whether a mod length == 0, because then we will get constant indexes for all elements in slice
	if a%length == 0 {
		return newMatcher(length)
	}
	b := crand.MrandRange(1, 1000)
	return func(i int) int {
		return (a*i + b) % length
	}
}

func Shuffle[T any](slice []T) []T {
	l := len(slice)
	result := make([]T, l)
	matcher := newMatcher(l)

	for i := range slice {
		result[matcher(i)] = slice[i]
	}
	return result
}
