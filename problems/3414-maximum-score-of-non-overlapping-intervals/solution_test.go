package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  []int
	}{
		{
			name:      "example 1: two intervals beat any triple",
			intervals: [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}, {6, 9, 3}, {6, 7, 1}, {8, 9, 1}},
			expected:  []int{2, 3},
		},
		{
			name:      "example 2: full budget of four intervals",
			intervals: [][]int{{5, 8, 1}, {6, 7, 7}, {4, 7, 3}, {9, 10, 6}, {7, 8, 2}, {11, 14, 3}, {3, 5, 5}},
			expected:  []int{1, 3, 5, 6},
		},
		{
			name:      "edge case: single interval must be chosen",
			intervals: [][]int{{1, 1000000000, 1000000000}},
			expected:  []int{0},
		},
		{
			name:      "edge case: all intervals mutually overlap, heaviest wins",
			intervals: [][]int{{1, 5, 3}, {2, 6, 4}, {3, 7, 3}},
			expected:  []int{1},
		},
		{
			name:      "edge case: shared boundary counts as overlapping",
			intervals: [][]int{{1, 5, 5}, {5, 10, 5}},
			expected:  []int{0},
		},
		{
			name:      "edge case: more than four disjoint intervals, keep heaviest four",
			intervals: [][]int{{1, 1, 1}, {3, 3, 2}, {5, 5, 3}, {7, 7, 4}, {9, 9, 5}},
			expected:  []int{1, 2, 3, 4},
		},
		{
			name:      "edge case: equal weights, smallest indices win",
			intervals: [][]int{{1, 1, 1}, {3, 3, 1}, {5, 5, 1}, {7, 7, 1}, {9, 9, 1}},
			expected:  []int{0, 1, 2, 3},
		},
		{
			name:      "edge case: tie between a pair and a single, pair has smaller indices",
			intervals: [][]int{{1, 4, 5}, {5, 10, 5}, {1, 10, 10}},
			expected:  []int{0, 1},
		},
		{
			name:      "edge case: tie between a single and a pair, single has smaller index",
			intervals: [][]int{{1, 10, 10}, {1, 4, 5}, {5, 10, 5}},
			expected:  []int{0},
		},
		{
			name:      "edge case: duplicate intervals, first index wins",
			intervals: [][]int{{1, 2, 5}, {1, 2, 5}},
			expected:  []int{0},
		},
		{
			name:      "edge case: maximal weights do not overflow the score",
			intervals: [][]int{{1, 2, 1000000000}, {3, 4, 1000000000}, {5, 6, 1000000000}, {7, 8, 1000000000}, {9, 10, 1000000000}},
			expected:  []int{0, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumWeight(tt.intervals)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("maximumWeight(%v) = %v, want %v", tt.intervals, result, tt.expected)
			}
		})
	}
}

// TestSolutionAgainstBruteForce cross-checks the DP against exhaustive search on
// small random inputs with a tiny coordinate and weight range, so that ties (and
// therefore the lexicographic tie-break) come up often.
func TestSolutionAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(42))

	for _, n := range []int{1, 2, 3, 5, 7} {
		for iter := 0; iter < 300; iter++ {
			intervals := make([][]int, n)
			for i := range intervals {
				l := rng.Intn(8) + 1
				r := l + rng.Intn(4)
				intervals[i] = []int{l, r, rng.Intn(3) + 1}
			}

			got := maximumWeight(intervals)
			want := bruteForceMaximumWeight(intervals)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("maximumWeight(%v) = %v, want %v", intervals, got, want)
			}
		}
	}
}

// bruteForceMaximumWeight tries every subset of at most 4 pairwise
// non-overlapping intervals.
func bruteForceMaximumWeight(intervals [][]int) []int {
	n := len(intervals)
	var bestScore int
	var best []int

	for mask := 1; mask < 1<<n; mask++ {
		var picked []int
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				picked = append(picked, i)
			}
		}
		if len(picked) > 4 {
			continue
		}

		score := 0
		ok := true
		for a := 0; a < len(picked) && ok; a++ {
			score += intervals[picked[a]][2]
			for b := a + 1; b < len(picked); b++ {
				x, y := intervals[picked[a]], intervals[picked[b]]
				if x[1] >= y[0] && y[1] >= x[0] {
					ok = false
					break
				}
			}
		}
		if !ok {
			continue
		}

		if best == nil || score > bestScore || (score == bestScore && lexLess(picked, best)) {
			bestScore, best = score, picked
		}
	}
	return best
}

func lexLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
