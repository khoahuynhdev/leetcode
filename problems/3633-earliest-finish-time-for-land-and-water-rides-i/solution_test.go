package main

import "testing"

func TestEarliestFinishTime(t *testing.T) {
	tests := []struct {
		name           string
		landStartTime  []int
		landDuration   []int
		waterStartTime []int
		waterDuration  []int
		expected       int
	}{
		{
			name:           "example 1: multiple land rides, single water ride",
			landStartTime:  []int{2, 8},
			landDuration:   []int{4, 1},
			waterStartTime: []int{6},
			waterDuration:  []int{3},
			expected:       9,
		},
		{
			name:           "example 2: water opens early but is long",
			landStartTime:  []int{5},
			landDuration:   []int{3},
			waterStartTime: []int{1},
			waterDuration:  []int{10},
			expected:       14,
		},
		{
			name:           "edge case: both rides open at time 1 with duration 1",
			landStartTime:  []int{1},
			landDuration:   []int{1},
			waterStartTime: []int{1},
			waterDuration:  []int{1},
			expected:       3,
		},
		{
			name:           "edge case: land much later than water finishes, water-then-land wins",
			landStartTime:  []int{100},
			landDuration:   []int{1},
			waterStartTime: []int{1},
			waterDuration:  []int{1},
			expected:       101,
		},
		{
			name:           "edge case: multiple rides on both sides, mixed durations",
			landStartTime:  []int{1, 10, 20},
			landDuration:   []int{5, 1, 1},
			waterStartTime: []int{2, 15},
			waterDuration:  []int{4, 1},
			expected:       10,
		},
		{
			name:           "edge case: max constraints (start=1000, duration=1000)",
			landStartTime:  []int{1000},
			landDuration:   []int{1000},
			waterStartTime: []int{1000},
			waterDuration:  []int{1000},
			expected:       3000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := earliestFinishTime(tt.landStartTime, tt.landDuration, tt.waterStartTime, tt.waterDuration)
			if got != tt.expected {
				t.Errorf("earliestFinishTime(%v, %v, %v, %v) = %d, want %d",
					tt.landStartTime, tt.landDuration, tt.waterStartTime, tt.waterDuration, got, tt.expected)
			}
		})
	}
}
