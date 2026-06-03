package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name           string
		landStartTime  []int
		landDuration   []int
		waterStartTime []int
		waterDuration  []int
		expected       int
	}{
		{
			name:           "example 1: two land rides, one water ride",
			landStartTime:  []int{2, 8},
			landDuration:   []int{4, 1},
			waterStartTime: []int{6},
			waterDuration:  []int{3},
			expected:       9,
		},
		{
			name:           "example 2: water first beats land first",
			landStartTime:  []int{5},
			landDuration:   []int{3},
			waterStartTime: []int{1},
			waterDuration:  []int{10},
			expected:       14,
		},
		{
			name:           "edge case: single ride each, no waiting",
			landStartTime:  []int{1},
			landDuration:   []int{1},
			waterStartTime: []int{1},
			waterDuration:  []int{1},
			expected:       3,
		},
		{
			name:           "edge case: second ride opens far in the future forces a wait",
			landStartTime:  []int{1},
			landDuration:   []int{2},
			waterStartTime: []int{100},
			waterDuration:  []int{5},
			expected:       105,
		},
		{
			name:           "edge case: best land ride for ordering is not the lowest index",
			landStartTime:  []int{10, 1},
			landDuration:   []int{10, 1},
			waterStartTime: []int{1},
			waterDuration:  []int{1},
			expected:       3,
		},
		{
			name:           "edge case: multiple rides on both sides",
			landStartTime:  []int{3, 7},
			landDuration:   []int{2, 1},
			waterStartTime: []int{4, 9},
			waterDuration:  []int{2, 1},
			expected:       7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := earliestFinishTime(tt.landStartTime, tt.landDuration, tt.waterStartTime, tt.waterDuration)
			if result != tt.expected {
				t.Errorf("earliestFinishTime(%v, %v, %v, %v) = %d, want %d",
					tt.landStartTime, tt.landDuration, tt.waterStartTime, tt.waterDuration, result, tt.expected)
			}
		})
	}
}
