package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		low      int
		high     int
		expected []int
	}{
		{
			name:     "example 1: range 100 to 300",
			low:      100,
			high:     300,
			expected: []int{123, 234},
		},
		{
			name:     "example 2: range 1000 to 13000",
			low:      1000,
			high:     13000,
			expected: []int{1234, 2345, 3456, 4567, 5678, 6789, 12345},
		},
		{
			name:     "edge case: no sequential digits in range",
			low:      140,
			high:     150,
			expected: []int{},
		},
		{
			name:     "edge case: single value range with one match",
			low:      123,
			high:     123,
			expected: []int{123},
		},
		{
			name:     "edge case: boundary spanning two- and three-digit numbers",
			low:      58,
			high:     155,
			expected: []int{67, 78, 89, 123},
		},
		{
			name:     "edge case: upper limit 10^9 has no 10-digit sequential number",
			low:      1000000000,
			high:     1000000000,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sequentialDigits(tt.low, tt.high)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sequentialDigits(%d, %d) = %v, want %v", tt.low, tt.high, result, tt.expected)
			}
		})
	}
}
