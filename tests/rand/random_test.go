package rand__test

import (
	"fmt"
	"github.com/foxclore/go-utils/crand"
	"testing"
)

func TestRandomSimple(t *testing.T) {
	n := crand.Random(10)
	fmt.Println(n)
}
