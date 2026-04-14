package main

import "testing"

func TestMinimumTotalDistanceDaily20260414(t *testing.T) {
	tests := []struct {
		name     string
		robot    []int
		factory  [][]int
		expected int64
	}{
		{
			name:     "example 1: three robots two factories",
			robot:    []int{0, 4, 6},
			factory:  [][]int{{2, 2}, {6, 2}},
			expected: 4,
		},
		{
			name:     "example 2: robots on opposite sides",
			robot:    []int{1, -1},
			factory:  [][]int{{-2, 1}, {2, 1}},
			expected: 2,
		},
		{
			name:     "edge case: single robot at factory position",
			robot:    []int{5},
			factory:  [][]int{{5, 1}},
			expected: 0,
		},
		{
			name:     "edge case: factory with limit zero must be skipped",
			robot:    []int{1},
			factory:  [][]int{{0, 0}, {2, 1}},
			expected: 1,
		},
		{
			name:     "edge case: unsorted inputs",
			robot:    []int{9, 0, -3, 5},
			factory:  [][]int{{10, 1}, {-5, 2}, {4, 2}},
			// Optimal: -3->-5 (2), 0->4 (4), 5->4 (1), 9->10 (1) = 8.
			expected: 8,
		},
		{
			name:     "edge case: large coordinates use int64",
			robot:    []int{-1000000000, 1000000000},
			factory:  [][]int{{0, 2}},
			expected: 2000000000,
		},
		{
			name:     "edge case: surplus capacity leaves factories unused",
			robot:    []int{0},
			factory:  [][]int{{-10, 3}, {10, 3}},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make copies so the solution's in-place sort doesn't affect later
			// assertions or sibling tests relying on original order.
			robotCopy := append([]int(nil), tt.robot...)
			factoryCopy := make([][]int, len(tt.factory))
			for i, f := range tt.factory {
				factoryCopy[i] = append([]int(nil), f...)
			}
			got := minimumTotalDistance(robotCopy, factoryCopy)
			if got != tt.expected {
				t.Errorf("minimumTotalDistance(%v, %v) = %d, want %d",
					tt.robot, tt.factory, got, tt.expected)
			}
		})
	}
}
