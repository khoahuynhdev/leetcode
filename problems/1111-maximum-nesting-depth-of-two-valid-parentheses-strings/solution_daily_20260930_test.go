package main

import "testing"

// splitDepth returns the depth of the subsequence of seq assigned to group g,
// or -1 if that subsequence is not a valid parentheses string.
func splitDepth(seq string, ans []int, g int) int {
	depth, best := 0, 0
	for i := 0; i < len(seq); i++ {
		if ans[i] != g {
			continue
		}
		if seq[i] == '(' {
			depth++
			if depth > best {
				best = depth
			}
		} else {
			depth--
			if depth < 0 {
				return -1
			}
		}
	}
	if depth != 0 {
		return -1
	}
	return best
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		seq      string
		maxDepth int // optimal max(depth(A), depth(B))
	}{
		{"example 1: nested pairs inside one pair", "(()())", 1},
		{"example 2: mix of flat and nested", "()(())()", 1},
		{"edge case: single pair", "()", 1},
		{"edge case: flat sequence", "()()()", 1},
		{"edge case: deep nesting even depth", "(((())))", 2},
		{"edge case: odd depth", "((()))", 2},
		{"edge case: depth five with siblings", "((((()))()))()", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := maxDepthAfterSplit(tt.seq)
			if len(ans) != len(tt.seq) {
				t.Fatalf("got length %d, want %d", len(ans), len(tt.seq))
			}
			for i, v := range ans {
				if v != 0 && v != 1 {
					t.Fatalf("ans[%d] = %d, want 0 or 1", i, v)
				}
			}
			a := splitDepth(tt.seq, ans, 0)
			b := splitDepth(tt.seq, ans, 1)
			if a < 0 || b < 0 {
				t.Fatalf("invalid split %v: group 0 depth %d, group 1 depth %d", ans, a, b)
			}
			got := max(a, b)
			if got != tt.maxDepth {
				t.Errorf("got max depth %d (split %v), want %d", got, ans, tt.maxDepth)
			}
		})
	}
}
