package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name          string
		n             int
		reservedSeats [][]int
		expected      int
	}{
		{
			name:          "example 1: mixed rows, one row untouched",
			n:             3,
			reservedSeats: [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}},
			expected:      4,
		},
		{
			name:          "example 2: every row blocked once",
			n:             2,
			reservedSeats: [][]int{{2, 1}, {1, 8}, {2, 6}},
			expected:      2,
		},
		{
			name:          "example 3: two rows fully blocked, two rows free",
			n:             4,
			reservedSeats: [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}},
			expected:      4,
		},
		{
			name:          "edge case: reservations only at seats 1 and 10 are harmless",
			n:             2,
			reservedSeats: [][]int{{1, 1}, {1, 10}, {2, 10}},
			expected:      4,
		},
		{
			name:          "edge case: single row where only the middle block survives",
			n:             1,
			reservedSeats: [][]int{{1, 2}, {1, 9}},
			expected:      1,
		},
		{
			name:          "edge case: single row with all three blocks broken",
			n:             1,
			reservedSeats: [][]int{{1, 5}, {1, 6}},
			expected:      0,
		},
		{
			name:          "edge case: huge n with a single reservation",
			n:             1000000000,
			reservedSeats: [][]int{{1, 5}},
			expected:      1999999999,
		},
		{
			name:          "edge case: seat 4 and 5 alone kill left and middle only",
			n:             1,
			reservedSeats: [][]int{{1, 4}, {1, 5}},
			expected:      1,
		},
		{
			name:          "edge case: one row untouched in the middle of blocked rows",
			n:             3,
			reservedSeats: [][]int{{1, 5}, {1, 6}, {3, 5}, {3, 6}},
			expected:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxNumberOfFamilies(tt.n, tt.reservedSeats)
			if result != tt.expected {
				t.Errorf("maxNumberOfFamilies(%d, %v) = %v, want %v", tt.n, tt.reservedSeats, result, tt.expected)
			}
		})
	}
}
