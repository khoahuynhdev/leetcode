package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		tasks    [][]int
		expected int
	}{
		{
			name:     "example 1: strictly increasing tasks",
			tasks:    [][]int{{1, 2}, {2, 4}, {4, 8}},
			expected: 8,
		},
		{
			name:     "example 2: mixed slacks",
			tasks:    [][]int{{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9}},
			expected: 32,
		},
		{
			name:     "example 3: constant slack across all tasks",
			tasks:    [][]int{{1, 7}, {2, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 12}},
			expected: 27,
		},
		{
			name:     "edge case: single task with a higher start requirement",
			tasks:    [][]int{{5, 10}},
			expected: 10,
		},
		{
			name:     "edge case: single task where actual equals minimum",
			tasks:    [][]int{{3, 3}},
			expected: 3,
		},
		{
			name:     "edge case: ordering between two tasks matters",
			tasks:    [][]int{{1, 2}, {1, 5}},
			expected: 5,
		},
		{
			name:     "edge case: all identical tasks",
			tasks:    [][]int{{2, 2}, {2, 2}, {2, 2}},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// minimumEffort sorts its argument in place; pass a copy so a
			// failure message can still report the original input.
			input := make([][]int, len(tt.tasks))
			for i, task := range tt.tasks {
				input[i] = append([]int(nil), task...)
			}
			if got := minimumEffort(input); got != tt.expected {
				t.Errorf("minimumEffort(%v) = %d, want %d", tt.tasks, got, tt.expected)
			}
		})
	}
}
