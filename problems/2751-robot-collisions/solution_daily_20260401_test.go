package main

import (
	"reflect"
	"testing"
)

func TestSurvivedRobotsHealths(t *testing.T) {
	tests := []struct {
		name       string
		positions  []int
		healths    []int
		directions string
		expected   []int
	}{
		{
			name:       "example 1: all moving right, no collisions",
			positions:  []int{5, 4, 3, 2, 1},
			healths:    []int{2, 17, 9, 15, 10},
			directions: "RRRRR",
			expected:   []int{2, 17, 9, 15, 10},
		},
		{
			name:       "example 2: two collisions, one survivor",
			positions:  []int{3, 5, 2, 6},
			healths:    []int{10, 10, 15, 12},
			directions: "RLRL",
			expected:   []int{14},
		},
		{
			name:       "example 3: all destroyed in equal collisions",
			positions:  []int{1, 2, 5, 6},
			healths:    []int{10, 10, 11, 11},
			directions: "RLRL",
			expected:   []int{},
		},
		{
			name:       "edge case: single robot",
			positions:  []int{42},
			healths:    []int{100},
			directions: "R",
			expected:   []int{100},
		},
		{
			name:       "edge case: all moving left, no collisions",
			positions:  []int{1, 2, 3, 4, 5},
			healths:    []int{5, 4, 3, 2, 1},
			directions: "LLLLL",
			expected:   []int{5, 4, 3, 2, 1},
		},
		{
			name:       "edge case: chain collision, high-health left-mover destroys many",
			positions:  []int{1, 2, 3, 4, 5},
			healths:    []int{1, 1, 1, 1, 100},
			directions: "RRRRL",
			expected:   []int{96},
		},
		{
			name:       "edge case: right-mover survives multiple left-movers",
			positions:  []int{1, 2, 3, 4, 5},
			healths:    []int{100, 1, 1, 1, 1},
			directions: "RLLLL",
			expected:   []int{96},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy inputs since the function modifies healths in place
			h := make([]int, len(tt.healths))
			copy(h, tt.healths)

			result := survivedRobotsHealths(tt.positions, h, tt.directions)

			if len(result) == 0 && len(tt.expected) == 0 {
				return // both empty, pass
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
