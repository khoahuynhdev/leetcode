package main

import "testing"

func TestMinNumberOfSeconds(t *testing.T) {
	tests := []struct {
		name           string
		mountainHeight int
		workerTimes    []int
		expected       int64
	}{
		{
			name:           "example 1: three workers mixed rates",
			mountainHeight: 4,
			workerTimes:    []int{2, 1, 1},
			expected:       3,
		},
		{
			name:           "example 2: four workers",
			mountainHeight: 10,
			workerTimes:    []int{3, 2, 2, 4},
			expected:       12,
		},
		{
			name:           "example 3: single worker",
			mountainHeight: 5,
			workerTimes:    []int{1},
			expected:       15,
		},
		{
			name:           "edge case: mountain height 1",
			mountainHeight: 1,
			workerTimes:    []int{5, 3, 7},
			expected:       3,
		},
		{
			name:           "edge case: single worker with large rate",
			mountainHeight: 3,
			workerTimes:    []int{1000000},
			expected:       6000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minNumberOfSeconds(tt.mountainHeight, tt.workerTimes)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
