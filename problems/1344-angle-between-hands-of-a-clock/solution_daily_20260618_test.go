package main

import (
	"math"
	"testing"
)

func TestAngleClock(t *testing.T) {
	const epsilon = 1e-5

	tests := []struct {
		name     string
		hour     int
		minutes  int
		expected float64
	}{
		{"example 1: 12:30 -> 165", 12, 30, 165},
		{"example 2: 3:30 -> 75", 3, 30, 75},
		{"example 3: 3:15 -> 7.5", 3, 15, 7.5},
		{"edge case: 12:00 hands overlap -> 0", 12, 0, 0},
		{"edge case: 6:00 hands opposite -> 180", 6, 0, 180},
		{"edge case: 3:00 right angle -> 90", 3, 0, 90},
		{"edge case: 1:00 -> 30", 1, 0, 30},
		{"edge case: 12:59 reflex handled -> 35.5", 12, 59, 35.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := angleClock(tt.hour, tt.minutes)
			if math.Abs(result-tt.expected) > epsilon {
				t.Errorf("angleClock(%d, %d) = %v, want %v", tt.hour, tt.minutes, result, tt.expected)
			}
		})
	}
}
