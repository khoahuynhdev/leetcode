package main

import (
	"math"
	"testing"
)

func TestChampagneTowerDaily(t *testing.T) {
	tests := []struct {
		name       string
		poured     int
		queryRow   int
		queryGlass int
		expected   float64
	}{
		{
			name:       "example 1: poured 1 cup, query (1,1)",
			poured:     1,
			queryRow:   1,
			queryGlass: 1,
			expected:   0.00000,
		},
		{
			name:       "example 2: poured 2 cups, query (1,1)",
			poured:     2,
			queryRow:   1,
			queryGlass: 1,
			expected:   0.50000,
		},
		{
			name:       "example 3: large pour, deep row",
			poured:     100000009,
			queryRow:   33,
			queryGlass: 17,
			expected:   1.00000,
		},
		{
			name:       "edge case: no champagne poured",
			poured:     0,
			queryRow:   0,
			queryGlass: 0,
			expected:   0.00000,
		},
		{
			name:       "edge case: top glass exactly full",
			poured:     1,
			queryRow:   0,
			queryGlass: 0,
			expected:   1.00000,
		},
		{
			name:       "edge case: 4 cups, third row middle glass",
			poured:     4,
			queryRow:   2,
			queryGlass: 1,
			expected:   0.50000,
		},
		{
			name:       "edge case: 4 cups, third row edge glass",
			poured:     4,
			queryRow:   2,
			queryGlass: 0,
			expected:   0.25000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := champagneTowerDaily(tt.poured, tt.queryRow, tt.queryGlass)
			if math.Abs(result-tt.expected) > 1e-5 {
				t.Errorf("champagneTowerDaily(%d, %d, %d) = %f, want %f",
					tt.poured, tt.queryRow, tt.queryGlass, result, tt.expected)
			}
		})
	}
}
