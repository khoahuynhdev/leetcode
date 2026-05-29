package main

import "testing"

func TestMinElementDaily20260529(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1: mixed two-digit numbers", []int{10, 12, 13, 14}, 1},
		{"example 2: already single digits", []int{1, 2, 3, 4}, 1},
		{"example 3: three-digit numbers", []int{999, 19, 199}, 10},
		{"edge case: single element", []int{7}, 7},
		{"edge case: maximum constraint value", []int{10000}, 1},
		{"edge case: all identical elements", []int{55, 55, 55}, 10},
		{"edge case: minimum hides behind large input", []int{9999, 9999, 100, 9999}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minElementDaily20260529(tt.nums)
			if got != tt.expected {
				t.Errorf("minElementDaily20260529(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
