package rand__test

import (
	"fmt"
	"github.com/foxclore/go-utils/v1/crand"
	"testing"
)

func TestRandomSimple(t *testing.T) {
	n := crand.Random(10)
	fmt.Println(n)
}
