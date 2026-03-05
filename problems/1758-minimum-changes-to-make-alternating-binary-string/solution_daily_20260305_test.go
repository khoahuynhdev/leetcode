package main

import "testing"

func TestMinOperationsDaily20260305(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"example 1: change last char", "0100", 1},
		{"example 2: already alternating", "10", 0},
		{"example 3: all ones", "1111", 2},
		{"edge case: single character 0", "0", 0},
		{"edge case: single character 1", "1", 0},
		{"edge case: all zeros even length", "0000", 2},
		{"edge case: already alternating starts with 0", "0101", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minOperationsDaily20260305(tt.input)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
