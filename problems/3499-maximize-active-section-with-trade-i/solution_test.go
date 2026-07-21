package main

import "testing"

func TestMaxActiveSectionsAfterTrade(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{"example 1: no valid trade", "01", 1},
		{"example 2: single one surrounded by zeros", "0100", 4},
		{"example 3: pick the wider sandwich", "1000100", 7},
		{"example 4: multiple equal sandwiches", "01010", 4},
		{"edge case: all ones", "111", 3},
		{"edge case: all zeros", "0000", 0},
		{"edge case: single zero", "0", 0},
		{"edge case: single one", "1", 1},
		{"edge case: no sandwich, ones at ends", "1001", 2},
		{"edge case: long tail sandwich", "10001000001", 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxActiveSectionsAfterTrade(tt.s); got != tt.expected {
				t.Errorf("maxActiveSectionsAfterTrade(%q) = %d, want %d", tt.s, got, tt.expected)
			}
		})
	}
}
