package main

import "testing"

func TestNumberOfStableArrays(t *testing.T) {
	tests := []struct {
		name     string
		zero     int
		one      int
		limit    int
		expected int
	}{
		{
			name:     "example 1: zero=1 one=1 limit=2",
			zero:     1,
			one:      1,
			limit:    2,
			expected: 2,
		},
		{
			name:     "example 2: zero=1 one=2 limit=1",
			zero:     1,
			one:      2,
			limit:    1,
			expected: 1,
		},
		{
			name:     "example 3: zero=3 one=3 limit=2",
			zero:     3,
			one:      3,
			limit:    2,
			expected: 14,
		},
		{
			name:     "edge case: all zeros within limit",
			zero:     5,
			one:      0,
			limit:    5,
			expected: 1,
		},
		{
			name:     "edge case: all zeros exceed limit",
			zero:     5,
			one:      0,
			limit:    4,
			expected: 0,
		},
		{
			name:     "edge case: limit=1 alternating possible",
			zero:     3,
			one:      3,
			limit:    1,
			expected: 2,
		},
		{
			name:     "edge case: limit=1 alternating impossible",
			zero:     3,
			one:      1,
			limit:    1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numberOfStableArrays(tt.zero, tt.one, tt.limit)
			if result != tt.expected {
				t.Errorf("numberOfStableArrays(%d, %d, %d) = %d, want %d",
					tt.zero, tt.one, tt.limit, result, tt.expected)
			}
		})
	}
}
