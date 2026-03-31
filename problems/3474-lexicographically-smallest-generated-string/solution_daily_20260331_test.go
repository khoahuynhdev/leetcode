package main

import "testing"

func TestGenerateString(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected string
	}{
		{
			name:     "example 1: TFTF with ab",
			str1:     "TFTF",
			str2:     "ab",
			expected: "ababa",
		},
		{
			name:     "example 2: TFTF with abc impossible",
			str1:     "TFTF",
			str2:     "abc",
			expected: "",
		},
		{
			name:     "example 3: single F with d",
			str1:     "F",
			str2:     "d",
			expected: "a",
		},
		{
			name:     "edge case: overlapping T conflict",
			str1:     "TT",
			str2:     "ab",
			expected: "",
		},
		{
			name:     "edge case: all F single char str2",
			str1:     "FFF",
			str2:     "a",
			expected: "bbb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateString(tt.str1, tt.str2)
			if result != tt.expected {
				t.Errorf("generateString(%q, %q) = %q, want %q", tt.str1, tt.str2, result, tt.expected)
			}
		})
	}
}
