package main

import "testing"

func TestDigitSumDaily20261116(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected string
	}{
		{"example 1: two rounds with multi-digit sum", "11111222223", 3, "135"},
		{"example 2: all zeros", "00000000", 3, "000"},
		{"edge case: length equals k, no rounds", "1234", 4, "1234"},
		{"edge case: length less than k", "7", 2, "7"},
		{"edge case: round keeps length then shrinks", "999", 2, "99"},
		{"edge case: multiple rounds, exact groups", "123456", 2, "12"},
		{"edge case: large k equals max length", "9999999999", 100, "9999999999"},
		{"edge case: many rounds with k=2", "99999999", 2, "99"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := digitSum(tt.s, tt.k)
			if got != tt.expected {
				t.Errorf("digitSum(%q, %d) = %q, want %q", tt.s, tt.k, got, tt.expected)
			}
		})
	}
}
