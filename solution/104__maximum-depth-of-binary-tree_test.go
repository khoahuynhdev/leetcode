package solution

import (
	"testing"
)

func Test104(t *testing.T) {
	if val := maxDepth(nil); val != 1 {
		t.Errorf("got %v", val)
	}
}
