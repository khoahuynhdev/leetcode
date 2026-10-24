package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name        string
		arriveAlice string
		leaveAlice  string
		arriveBob   string
		leaveBob    string
		expected    int
	}{
		{"example 1: partial overlap in same month", "08-15", "08-18", "08-16", "08-19", 3},
		{"example 2: no overlap", "10-01", "10-31", "11-01", "12-31", 0},
		{"edge case: touching on a single day", "01-01", "03-15", "03-15", "05-01", 1},
		{"edge case: identical single day", "06-30", "06-30", "06-30", "06-30", 1},
		{"edge case: one interval inside the other", "01-01", "12-31", "02-27", "03-02", 4},
		{"edge case: whole year together", "01-01", "12-31", "01-01", "12-31", 365},
		{"edge case: Bob before Alice, no overlap", "09-01", "09-10", "02-01", "08-31", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countDaysTogether(tt.arriveAlice, tt.leaveAlice, tt.arriveBob, tt.leaveBob)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
