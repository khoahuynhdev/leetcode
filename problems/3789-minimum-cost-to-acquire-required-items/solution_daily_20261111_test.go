package main

import "testing"

func TestMinimumCost(t *testing.T) {
	tests := []struct {
		name     string
		cost1    int
		cost2    int
		costBoth int
		need1    int
		need2    int
		expected int64
	}{
		{"example 1: both items cheapest, buy only type 3", 3, 2, 1, 3, 2, 3},
		{"example 2: both item too expensive, buy singles", 5, 4, 15, 2, 3, 22},
		{"example 3: nothing needed", 5, 4, 15, 0, 0, 0},
		{"edge case: both cheaper than pair but not single, overlap only", 5, 4, 7, 2, 5, 26},
		{"edge case: one need zero, both cheaper than single", 10, 10, 3, 0, 4, 12},
		{"edge case: one need zero, single cheaper", 10, 2, 3, 0, 4, 8},
		{"edge case: costBoth equals sum of singles", 3, 4, 7, 5, 5, 35},
		{"edge case: max constraints all type 3", 1000000, 1000000, 1, 1000000000, 1000000000, 1000000000},
		{"edge case: max constraints large result", 1000000, 1000000, 1000000, 1000000000, 1000000000, 1000000000000000},
		{"edge case: max constraints singles only", 1000000, 1000000, 1000000, 1000000000, 0, 1000000000000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumCost(tt.cost1, tt.cost2, tt.costBoth, tt.need1, tt.need2)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
