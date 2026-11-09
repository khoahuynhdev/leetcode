package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestBusiestServers(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		arrival  []int
		load     []int
		expected []int
	}{
		{
			name:     "example 1: request wraps to next free server, last one dropped",
			k:        3,
			arrival:  []int{1, 2, 3, 4, 5},
			load:     []int{5, 2, 3, 3, 3},
			expected: []int{1},
		},
		{
			name:     "example 2: preferred server is free again",
			k:        3,
			arrival:  []int{1, 2, 3, 4},
			load:     []int{1, 2, 1, 2},
			expected: []int{0},
		},
		{
			name:     "example 3: each server handles one request",
			k:        3,
			arrival:  []int{1, 2, 3},
			load:     []int{10, 12, 11},
			expected: []int{0, 1, 2},
		},
		{
			name:     "edge case: single server handles all non-overlapping requests",
			k:        1,
			arrival:  []int{1, 2, 3},
			load:     []int{1, 1, 1},
			expected: []int{0},
		},
		{
			name:     "edge case: single server drops overlapping requests",
			k:        1,
			arrival:  []int{1, 2, 3},
			load:     []int{5, 1, 1},
			expected: []int{0},
		},
		{
			name:     "edge case: more servers than requests",
			k:        5,
			arrival:  []int{1, 2},
			load:     []int{1, 1},
			expected: []int{0, 1},
		},
		{
			name:     "edge case: wrap-around past last server to server 0",
			k:        3,
			arrival:  []int{1, 2, 3, 4},
			load:     []int{1, 10, 10, 1},
			expected: []int{0},
		},
		{
			name:     "edge case: server freed exactly at arrival time is reusable",
			k:        2,
			arrival:  []int{1, 2, 3},
			load:     []int{2, 100, 1},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := busiestServers(tt.k, tt.arrival, tt.load)
			sort.Ints(result)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
