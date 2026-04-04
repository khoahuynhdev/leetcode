package main

import "testing"

func TestDecodeCiphertext(t *testing.T) {
	tests := []struct {
		name        string
		encodedText string
		rows        int
		expected    string
	}{
		{
			name:        "example 1: cipher",
			encodedText: "ch   ie   pr",
			rows:        3,
			expected:    "cipher",
		},
		{
			name:        "example 2: i love leetcode",
			encodedText: "iveo    eed   l te   olc",
			rows:        4,
			expected:    "i love leetcode",
		},
		{
			name:        "example 3: single row",
			encodedText: "coding",
			rows:        1,
			expected:    "coding",
		},
		{
			name:        "edge case: empty string",
			encodedText: "",
			rows:        5,
			expected:    "",
		},
		{
			name:        "edge case: single character",
			encodedText: "a",
			rows:        1,
			expected:    "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := decodeCiphertext(tt.encodedText, tt.rows)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
