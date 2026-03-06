package main

import "testing"

func TestCheckOnesSegment(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"example 1: ones split by zeros", "1001", false},
		{"example 2: contiguous ones then zero", "110", true},
		{"edge case: single character", "1", true},
		{"edge case: all ones", "1111", true},
		{"edge case: ones then zeros", "1100", true},
		{"edge case: alternating", "101", false},
		{"edge case: one followed by zero", "10", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkOnesSegment(tt.input)
			if result != tt.expected {
				t.Errorf("checkOnesSegment(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
