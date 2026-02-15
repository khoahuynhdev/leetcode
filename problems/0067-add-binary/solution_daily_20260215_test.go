package main

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{"example 1: simple carry", "11", "1", "100"},
		{"example 2: same length addition", "1010", "1011", "10101"},
		{"edge case: both zeros", "0", "0", "0"},
		{"edge case: one is zero", "0", "1", "1"},
		{"edge case: single ones produce carry", "1", "1", "10"},
		{"different lengths", "1", "11111", "100000"},
		{"all ones cascade carry", "1111", "1111", "11110"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addBinary(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("addBinary(%q, %q) = %q, want %q", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
