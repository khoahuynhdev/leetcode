package main

import (
	"reflect"
	"testing"
)

func TestAssignTasks(t *testing.T) {
	tests := []struct {
		name     string
		servers  []int
		tasks    []int
		expected []int
	}{
		{
			name:     "example 1: tie on weight broken by index",
			servers:  []int{3, 3, 2},
			tasks:    []int{1, 2, 3, 2, 1, 2},
			expected: []int{2, 2, 0, 2, 1, 2},
		},
		{
			name:     "example 2: servers freed and reused",
			servers:  []int{5, 1, 4, 3, 2},
			tasks:    []int{2, 1, 2, 4, 5, 2, 1},
			expected: []int{1, 4, 1, 4, 1, 3, 2},
		},
		{
			name:     "edge case: single server single task",
			servers:  []int{7},
			tasks:    []int{3},
			expected: []int{0},
		},
		{
			name:     "edge case: single server, tasks must wait in queue",
			servers:  []int{1},
			tasks:    []int{5, 1, 1},
			expected: []int{0, 0, 0},
		},
		{
			name:     "edge case: all busy, multiple servers free simultaneously",
			servers:  []int{2, 1},
			tasks:    []int{10, 10, 1, 1},
			expected: []int{1, 0, 1, 1},
		},
		{
			name:     "edge case: identical weights pick smallest index",
			servers:  []int{1, 1, 1},
			tasks:    []int{1, 1, 1},
			expected: []int{0, 0, 0},
		},
		{
			name:     "edge case: more servers than tasks, long tasks",
			servers:  []int{4, 3, 2, 1},
			tasks:    []int{100, 100},
			expected: []int{3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := assignTasks(tt.servers, tt.tasks)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("assignTasks(%v, %v) = %v, want %v", tt.servers, tt.tasks, got, tt.expected)
			}
		})
	}
}
