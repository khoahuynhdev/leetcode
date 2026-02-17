package main

import (
	"sort"
	"testing"
)

func TestReadBinaryWatch(t *testing.T) {
	tests := []struct {
		name     string
		turnedOn int
		expected []string
	}{
		{
			name:     "example 1: turnedOn=1",
			turnedOn: 1,
			expected: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{
			name:     "example 2: turnedOn=9 impossible",
			turnedOn: 9,
			expected: []string{},
		},
		{
			name:     "edge case: turnedOn=0 only midnight",
			turnedOn: 0,
			expected: []string{"0:00"},
		},
		{
			name:     "edge case: turnedOn=10 impossible",
			turnedOn: 10,
			expected: []string{},
		},
		{
			name:     "turnedOn=2 multiple combinations",
			turnedOn: 2,
			expected: []string{
				"0:03", "0:05", "0:06", "0:09", "0:10", "0:12", "0:17", "0:18", "0:20", "0:24",
				"0:33", "0:34", "0:36", "0:40", "0:48",
				"1:01", "1:02", "1:04", "1:08", "1:16", "1:32",
				"2:01", "2:02", "2:04", "2:08", "2:16", "2:32",
				"3:00",
				"4:01", "4:02", "4:04", "4:08", "4:16", "4:32",
				"5:00", "6:00",
				"8:01", "8:02", "8:04", "8:08", "8:16", "8:32",
				"9:00", "10:00",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := readBinaryWatch(tt.turnedOn)

			// Sort both slices for comparison since order doesn't matter
			sort.Strings(result)
			sort.Strings(tt.expected)

			if len(tt.expected) == 0 && len(result) == 0 {
				return // both empty, pass
			}

			if len(result) != len(tt.expected) {
				t.Errorf("got %d results, want %d\ngot:  %v\nwant: %v", len(result), len(tt.expected), result, tt.expected)
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("mismatch at index %d: got %q, want %q\nfull got:  %v\nfull want: %v", i, result[i], tt.expected[i], result, tt.expected)
					return
				}
			}
		})
	}
}
