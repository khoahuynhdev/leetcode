package main

import "testing"

func TestStoneGame(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		expected bool
	}{
		{"example 1: [5,3,4,5] Alice opens with the first 5", []int{5, 3, 4, 5}, true},
		{"example 2: [3,7,2,3]", []int{3, 7, 2, 3}, true},
		{"edge case: empty input, nobody wins", nil, false},
		{"edge case: minimum size, two unequal piles", []int{1, 2}, true},
		{"edge case: minimum size, larger pile first", []int{2, 1}, true},
		{"edge case: two equal piles tie, not a win", []int{1, 1}, false},
		{"edge case: odd pile count where Alice loses", []int{1, 5, 2}, false},
		{"edge case: odd pile count where Alice wins", []int{3, 1, 5}, true},
		{"edge case: strictly increasing piles", []int{1, 2, 3, 4, 5, 6}, true},
		{"edge case: strictly decreasing piles", []int{6, 5, 4, 3, 2, 1}, true},
		{"edge case: big middle pile is a trap", []int{2, 100, 99, 2}, true},
		{"edge case: max-value piles at the bounds", []int{500, 1, 2, 500}, true},
		{"edge case: even total makes the game a draw", []int{500, 1, 1, 500}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGame(tt.piles); got != tt.expected {
				t.Errorf("stoneGame(%v) = %v, want %v", tt.piles, got, tt.expected)
			}
		})
	}
}

// The input slice must not be mutated: LeetCode reuses caller-owned data.
func TestStoneGameDoesNotMutateInput(t *testing.T) {
	piles := []int{5, 3, 4, 5}
	original := append([]int(nil), piles...)

	stoneGame(piles)

	for i := range piles {
		if piles[i] != original[i] {
			t.Fatalf("input mutated: got %v, want %v", piles, original)
		}
	}
}

// Any valid input (even pile count, odd total) is a win for Alice.
func TestStoneGameAlwaysTrueForValidInputs(t *testing.T) {
	inputs := [][]int{
		{1, 2},
		{5, 3, 4, 5},
		{3, 7, 2, 3},
		{1, 1, 1, 2},
		{500, 500, 500, 499},
		{2, 4, 6, 8, 10, 500, 4, 1},
	}

	for _, piles := range inputs {
		sum := 0
		for _, p := range piles {
			sum += p
		}
		if len(piles)%2 != 0 || sum%2 == 0 {
			t.Fatalf("test input %v violates the problem constraints", piles)
		}
		if !stoneGame(piles) {
			t.Errorf("stoneGame(%v) = false, want true", piles)
		}
	}
}

func BenchmarkStoneGame(b *testing.B) {
	piles := make([]int, 500)
	for i := range piles {
		piles[i] = (i*37)%500 + 1
	}
	piles[0]++ // keep the total odd

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stoneGame(piles)
	}
}
