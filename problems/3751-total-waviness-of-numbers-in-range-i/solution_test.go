package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{"example 1: range [120, 130]", 120, 130, 3},
		{"example 2: range [198, 202]", 198, 202, 3},
		{"example 3: single number 4848 has peak and valley", 4848, 4848, 2},
		{"edge case: all numbers fewer than 3 digits", 1, 99, 0},
		{"edge case: single 2-digit number has no interior", 42, 42, 0},
		{"edge case: plateau equal neighbors do not count (121 vs 122)", 122, 122, 0},
		{"edge case: 121 middle is a peak", 121, 121, 1},
		{"edge case: 2002 equal interior neighbors count nothing", 2002, 2002, 0},
		{"edge case: 1213 has peak then valley", 1213, 1213, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalWaviness(tt.num1, tt.num2); got != tt.expected {
				t.Errorf("totalWaviness(%d, %d) = %d, want %d", tt.num1, tt.num2, got, tt.expected)
			}
		})
	}
}
