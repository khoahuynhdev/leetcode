package main

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		numRows  int
		expected string
	}{
		{
			name:     "example 1: PAYPALISHIRING with 3 rows",
			s:        "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "example 2: PAYPALISHIRING with 4 rows",
			s:        "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			name:     "example 3: single character",
			s:        "A",
			numRows:  1,
			expected: "A",
		},
		{
			name:     "edge case: numRows is 1",
			s:        "ABCDEFGH",
			numRows:  1,
			expected: "ABCDEFGH",
		},
		{
			name:     "edge case: string length equals numRows",
			s:        "ABC",
			numRows:  3,
			expected: "ABC",
		},
		{
			name:     "edge case: string shorter than numRows",
			s:        "AB",
			numRows:  5,
			expected: "AB",
		},
		{
			name:     "edge case: two rows",
			s:        "ABCDE",
			numRows:  2,
			expected: "ACEBD",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convert(tt.s, tt.numRows)
			if result != tt.expected {
				t.Errorf("convert(%q, %d) = %q, want %q", tt.s, tt.numRows, result, tt.expected)
			}
		})
	}
}
