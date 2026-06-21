package main

import "testing"

func TestMaxIceCream(t *testing.T) {
	tests := []struct {
		name     string
		costs    []int
		coins    int
		expected int
	}{
		{"example 1: buy four cheapest bars", []int{1, 3, 2, 4, 1}, 7, 4},
		{"example 2: cannot afford any bar", []int{10, 6, 8, 7, 7, 8}, 5, 0},
		{"example 3: afford all bars", []int{1, 6, 3, 1, 2, 5}, 20, 6},
		{"edge case: empty input", []int{}, 100, 0},
		{"edge case: single bar affordable", []int{5}, 5, 1},
		{"edge case: single bar unaffordable", []int{5}, 4, 0},
		{"edge case: all same price, partial buy", []int{2, 2, 2, 2}, 5, 2},
		{"edge case: exact budget for all", []int{3, 3, 3}, 9, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxIceCream(tt.costs, tt.coins)
			if result != tt.expected {
				t.Errorf("maxIceCream(%v, %d) = %d, want %d", tt.costs, tt.coins, result, tt.expected)
			}
		})
	}
}
