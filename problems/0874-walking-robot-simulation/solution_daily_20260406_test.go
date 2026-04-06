package main

import "testing"

func TestRobotSim(t *testing.T) {
	tests := []struct {
		name      string
		commands  []int
		obstacles [][]int
		expected  int
	}{
		{
			name:      "example 1: no obstacles",
			commands:  []int{4, -1, 3},
			obstacles: [][]int{},
			expected:  25,
		},
		{
			name:      "example 2: obstacle blocks east movement",
			commands:  []int{4, -1, 4, -2, 4},
			obstacles: [][]int{{2, 4}},
			expected:  65,
		},
		{
			name:      "example 3: obstacle at origin",
			commands:  []int{6, -1, -1, 6},
			obstacles: [][]int{{0, 0}},
			expected:  36,
		},
		{
			name:      "edge case: only turns, no movement",
			commands:  []int{-1, -2, -1, -2},
			obstacles: [][]int{},
			expected:  0,
		},
		{
			name:      "edge case: blocked on first step",
			commands:  []int{3},
			obstacles: [][]int{{0, 1}},
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := robotSim(tt.commands, tt.obstacles)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
