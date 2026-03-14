package main

import "testing"

func TestGetHappyString(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected string
	}{
		{"example 1: n=1 k=3 returns c", 1, 3, "c"},
		{"example 2: n=1 k=4 exceeds count returns empty", 1, 4, ""},
		{"example 3: n=3 k=9 returns cab", 3, 9, "cab"},
		{"edge case: first happy string", 1, 1, "a"},
		{"edge case: n=2 k=1 first of length 2", 2, 1, "ab"},
		{"edge case: k exceeds total for n=2", 2, 7, ""},
		{"edge case: last happy string of length 3", 3, 12, "cbc"},
		{"edge case: n=10 k=1 first long string", 10, 1, "ababababab"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getHappyString(tt.n, tt.k)
			if result != tt.expected {
				t.Errorf("getHappyString(%d, %d) = %q, want %q", tt.n, tt.k, result, tt.expected)
			}
		})
	}
}
