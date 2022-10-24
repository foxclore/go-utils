package everything

import (
	"github.com/foxclore/go-utils/everything"
	"testing"
)

func TestShuffle(t *testing.T) {
	l := []int{1, 2, 3, 4, 5}

	for i := 0; i < 10; i++ {
		t.Log(everything.Shuffle(l))
	}
}
