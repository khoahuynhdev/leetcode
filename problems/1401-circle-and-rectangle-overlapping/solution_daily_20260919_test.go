package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		radius   int
		xCenter  int
		yCenter  int
		x1       int
		y1       int
		x2       int
		y2       int
		expected bool
	}{
		{
			name:    "example 1: circle touches rectangle at the single point (1,0)",
			radius:  1,
			xCenter: 0, yCenter: 0,
			x1: 1, y1: -1, x2: 3, y2: 1,
			expected: true,
		},
		{
			name:    "example 2: center aligned horizontally but rectangle sits below",
			radius:  1,
			xCenter: 1, yCenter: 1,
			x1: 1, y1: -3, x2: 2, y2: -1,
			expected: false,
		},
		{
			name:    "example 3: rectangle corner coincides with the circle center",
			radius:  1,
			xCenter: 0, yCenter: 0,
			x1: -1, y1: 0, x2: 0, y2: 1,
			expected: true,
		},
		{
			name:    "edge case: center strictly inside the rectangle",
			radius:  1,
			xCenter: 5, yCenter: 5,
			x1: 0, y1: 0, x2: 10, y2: 10,
			expected: true,
		},
		{
			name:    "edge case: rectangle entirely contained in the circle",
			radius:  2000,
			xCenter: 0, yCenter: 0,
			x1: -1, y1: -1, x2: 1, y2: 1,
			expected: true,
		},
		{
			name:    "edge case: diagonally past a corner, both axis ranges overlap the circle",
			radius:  1,
			xCenter: 0, yCenter: 0,
			x1: 1, y1: 1, x2: 2, y2: 2,
			expected: false,
		},
		{
			name:    "edge case: exact tangency at a corner (3,4) with radius 5",
			radius:  5,
			xCenter: 0, yCenter: 0,
			x1: 3, y1: 4, x2: 10, y2: 10,
			expected: true,
		},
		{
			name:    "edge case: exact tangency against a vertical edge",
			radius:  2,
			xCenter: 0, yCenter: 0,
			x1: 2, y1: -5, x2: 5, y2: 5,
			expected: true,
		},
		{
			name:    "edge case: vertical edge one unit too far away",
			radius:  2,
			xCenter: 0, yCenter: 0,
			x1: 3, y1: -5, x2: 5, y2: 5,
			expected: false,
		},
		{
			name:    "edge case: extreme opposite corners of the coordinate range",
			radius:  2000,
			xCenter: -10000, yCenter: -10000,
			x1: 9999, y1: 9999, x2: 10000, y2: 10000,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkOverlap(tt.radius, tt.xCenter, tt.yCenter, tt.x1, tt.y1, tt.x2, tt.y2)
			if got != tt.expected {
				t.Errorf("checkOverlap(%d, %d, %d, %d, %d, %d, %d) = %v, want %v",
					tt.radius, tt.xCenter, tt.yCenter, tt.x1, tt.y1, tt.x2, tt.y2, got, tt.expected)
			}
		})
	}
}
