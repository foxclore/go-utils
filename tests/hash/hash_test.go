package hash__test

import (
	"fmt"
	"github.com/foxclore/go-utils/hash"
	"testing"
)

func TestHashSimple(t *testing.T) {
	h := hash.Hash(123, 123, 32453245, "234234")
	fmt.Println(h)
}
