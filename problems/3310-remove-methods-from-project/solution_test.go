package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestRemainingMethods(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		k           int
		invocations [][]int
		expected    []int
	}{
		{
			name:        "example 1: outside methods invoke the suspicious group, remove nothing",
			n:           4,
			k:           1,
			invocations: [][]int{{1, 2}, {0, 1}, {3, 2}},
			expected:    []int{0, 1, 2, 3},
		},
		{
			name:        "example 2: suspicious group is self-contained, remove it",
			n:           5,
			k:           0,
			invocations: [][]int{{1, 2}, {0, 2}, {0, 1}, {3, 4}},
			expected:    []int{3, 4},
		},
		{
			name:        "example 3: cycle makes every method suspicious",
			n:           3,
			k:           2,
			invocations: [][]int{{1, 2}, {0, 1}, {2, 0}},
			expected:    []int{},
		},
		{
			name:        "edge case: empty invocations, only k is suspicious",
			n:           4,
			k:           2,
			invocations: [][]int{},
			expected:    []int{0, 1, 3},
		},
		{
			name:        "edge case: single method with no invocations",
			n:           1,
			k:           0,
			invocations: [][]int{},
			expected:    []int{},
		},
		{
			name:        "edge case: blocking edge points deep into the suspicious chain",
			n:           6,
			k:           3,
			invocations: [][]int{{3, 4}, {4, 5}, {2, 5}},
			expected:    []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:        "edge case: suspicious chain is isolated from the rest",
			n:           6,
			k:           3,
			invocations: [][]int{{3, 4}, {4, 5}, {0, 1}, {1, 2}},
			expected:    []int{0, 1, 2},
		},
		{
			name:        "edge case: k invoked from outside blocks removal even with no outgoing edges",
			n:           3,
			k:           1,
			invocations: [][]int{{0, 1}},
			expected:    []int{0, 1, 2},
		},
		{
			name:        "edge case: disconnected node is neither suspicious nor a blocker",
			n:           3,
			k:           0,
			invocations: [][]int{{0, 1}},
			expected:    []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := remainingMethods(tt.n, tt.k, tt.invocations)
			// The problem allows any order, so normalize before comparing.
			sort.Ints(got)
			if len(got) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("remainingMethods(%d, %d, %v) = %v, want %v",
					tt.n, tt.k, tt.invocations, got, tt.expected)
			}
		})
	}
}
