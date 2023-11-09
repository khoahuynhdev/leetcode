package helper

import (
	"testing"
)

type argZZ struct {
	Base   int
	Pow    int
	Expect uint64
}

var table = []argZZ{
	{1, 1, 1},
	{2, 2, 4},
	{2, 3, 8},
}

func TestExponentiation(t *testing.T) {
	for _, test := range table {
		if res := FastExp(test.Base, test.Pow); res != test.Expect {
			t.Errorf("got %d; want %d for inputs = %d %d ", res, test.Expect, test.Base, test.Pow)
		}
	}
}
