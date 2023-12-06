package solution

import (
	"testing"
)

func Test104(t *testing.T) {
	if val := MaxDepth(nil); val != 1 {
		t.Errorf("got %v", val)
	}
}
