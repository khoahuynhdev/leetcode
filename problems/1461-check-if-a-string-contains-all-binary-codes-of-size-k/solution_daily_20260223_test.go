package main

import "testing"

func TestHasAllCodes(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected bool
	}{
		{
			name:     "example 1: all codes of length 2 present",
			s:        "00110110",
			k:        2,
			expected: true,
		},
		{
			name:     "example 2: all codes of length 1 present",
			s:        "0110",
			k:        1,
			expected: true,
		},
		{
			name:     "example 3: missing code 00 of length 2",
			s:        "0110",
			k:        2,
			expected: false,
		},
		{
			name:     "edge case: string shorter than k",
			s:        "01",
			k:        3,
			expected: false,
		},
		{
			name:     "edge case: k=1 with only zeros",
			s:        "0000",
			k:        1,
			expected: false,
		},
		{
			name:     "edge case: all same characters for k=2",
			s:        "0000",
			k:        2,
			expected: false,
		},
		{
			name:     "edge case: not enough windows for all codes",
			s:        "01",
			k:        2,
			expected: false,
		},
		{
			name:     "all codes of length 3 present",
			s:        "00010011101100101",
			k:        3,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasAllCodes(tt.s, tt.k)
			if result != tt.expected {
				t.Errorf("hasAllCodes(%q, %d) = %v, want %v", tt.s, tt.k, result, tt.expected)
			}
		})
	}
}
