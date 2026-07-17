package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		queries  []int64
		expected []int
	}{
		{
			name:     "example 1: [2,3,4] queries [0,2,2]",
			nums:     []int{2, 3, 4},
			queries:  []int64{0, 2, 2},
			expected: []int{1, 2, 2},
		},
		{
			name:     "example 2: [4,4,2,1] queries [5,3,1,0]",
			nums:     []int{4, 4, 2, 1},
			queries:  []int64{5, 3, 1, 0},
			expected: []int{4, 2, 1, 1},
		},
		{
			name:     "example 3: [2,2] queries [0,0]",
			nums:     []int{2, 2},
			queries:  []int64{0, 0},
			expected: []int{2, 2},
		},
		{
			name:     "edge case: single pair coprime",
			nums:     []int{3, 5},
			queries:  []int64{0},
			expected: []int{1},
		},
		{
			name:     "edge case: all duplicates yield same gcd",
			nums:     []int{6, 6, 6},
			queries:  []int64{0, 1, 2},
			expected: []int{6, 6, 6},
		},
		{
			name: "edge case: full sorted gcdPairs scanned in order",
			// nums = [4,4,2,1] -> gcdPairs sorted = [1,1,1,2,2,4]
			nums:     []int{4, 4, 2, 1},
			queries:  []int64{0, 1, 2, 3, 4, 5},
			expected: []int{1, 1, 1, 2, 2, 4},
		},
		{
			name:     "edge case: mixed values with larger gcds",
			// nums = [10,15,20] -> gcd(10,15)=5, gcd(10,20)=10, gcd(15,20)=5
			// sorted = [5,5,10]
			nums:     []int{10, 15, 20},
			queries:  []int64{0, 1, 2},
			expected: []int{5, 5, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gcdValues(tt.nums, tt.queries)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("gcdValues(%v, %v) = %v, want %v", tt.nums, tt.queries, result, tt.expected)
			}
		})
	}
}
