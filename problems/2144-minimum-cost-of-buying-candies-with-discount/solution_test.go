package main

import "testing"

func TestMinimumCost(t *testing.T) {
	tests := []struct {
		name     string
		cost     []int
		expected int
	}{
		{"example 1: three candies", []int{1, 2, 3}, 5},
		{"example 2: six candies with two free", []int{6, 5, 7, 9, 2, 2}, 23},
		{"example 3: two candies pay both", []int{5, 5}, 10},
		{"edge case: single candy", []int{7}, 7},
		{"edge case: four candies, one free", []int{1, 2, 3, 4}, 8},
		{"edge case: all equal costs", []int{5, 5, 5}, 10},
		{"edge case: max constraint length all 100s", func() []int {
			c := make([]int, 9)
			for i := range c {
				c[i] = 100
			}
			return c
		}(), 600},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.cost))
			copy(input, tt.cost)
			got := minimumCost(input)
			if got != tt.expected {
				t.Errorf("minimumCost(%v) = %d, want %d", tt.cost, got, tt.expected)
			}
		})
	}
}
