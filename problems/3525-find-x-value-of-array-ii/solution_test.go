package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		queries  [][]int
		expected []int
	}{
		{
			name:     "example 1: three queries with moving prefix start",
			nums:     []int{1, 2, 3, 4, 5},
			k:        3,
			queries:  [][]int{{2, 2, 0, 2}, {3, 3, 3, 0}, {0, 1, 0, 1}},
			expected: []int{2, 2, 2},
		},
		{
			name:     "example 2: same update twice, different x",
			nums:     []int{1, 2, 4, 8, 16, 32},
			k:        4,
			queries:  [][]int{{0, 2, 0, 2}, {0, 2, 0, 1}},
			expected: []int{1, 0},
		},
		{
			name:     "example 3: every prefix matches",
			nums:     []int{1, 1, 2, 1, 1},
			k:        2,
			queries:  [][]int{{2, 1, 0, 1}},
			expected: []int{5},
		},
		{
			name:     "edge case: single element array",
			nums:     []int{7},
			k:        3,
			queries:  [][]int{{0, 7, 0, 1}, {0, 6, 0, 0}, {0, 6, 0, 2}},
			expected: []int{1, 1, 0},
		},
		{
			name:     "edge case: k equals 1 so every product is 0",
			nums:     []int{1, 2, 3, 4},
			k:        1,
			queries:  [][]int{{0, 5, 0, 0}, {2, 9, 2, 0}},
			expected: []int{4, 2},
		},
		{
			name:     "edge case: start at the last index",
			nums:     []int{1, 2, 3},
			k:        3,
			queries:  [][]int{{1, 2, 2, 0}, {2, 4, 2, 1}},
			expected: []int{1, 1},
		},
		{
			name:     "edge case: a zero residue absorbs every later prefix",
			nums:     []int{2, 3, 4},
			k:        2,
			queries:  [][]int{{0, 3, 0, 1}, {0, 3, 0, 0}},
			expected: []int{2, 1},
		},
		{
			name:     "edge case: large values near the constraint limit",
			nums:     []int{1000000000, 999999999},
			k:        5,
			queries:  [][]int{{0, 1000000000, 0, 0}, {0, 999999999, 0, 4}},
			expected: []int{2, 1},
		},
		{
			name:     "edge case: updates persist across queries",
			nums:     []int{3, 4},
			k:        5,
			queries:  [][]int{{1, 5, 0, 0}, {0, 2, 1, 0}},
			expected: []int{1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resultArray(tt.nums, tt.k, tt.queries)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("resultArray(%v, %d, %v) = %v, want %v", tt.nums, tt.k, tt.queries, got, tt.expected)
			}
		})
	}
}

// bruteForce recomputes every prefix product for each query in O(n) time.
func bruteForce(nums []int, k int, queries [][]int) []int {
	arr := make([]int, len(nums))
	copy(arr, nums)

	res := make([]int, len(queries))
	for i, q := range queries {
		arr[q[0]] = q[1]
		cnt, prod := 0, 1%k
		for j := q[2]; j < len(arr); j++ {
			prod = prod * (arr[j] % k) % k
			if prod == q[3] {
				cnt++
			}
		}
		res[i] = cnt
	}
	return res
}

func TestSolutionAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(3525))

	for iter := 0; iter < 300; iter++ {
		n := 1 + rng.Intn(12)
		k := 1 + rng.Intn(5)

		nums := make([]int, n)
		for i := range nums {
			nums[i] = 1 + rng.Intn(20)
		}

		queries := make([][]int, 1+rng.Intn(8))
		for i := range queries {
			queries[i] = []int{rng.Intn(n), 1 + rng.Intn(20), rng.Intn(n), rng.Intn(k)}
		}

		want := bruteForce(nums, k, queries)
		got := resultArray(nums, k, queries)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("nums=%v k=%d queries=%v: got %v, want %v", nums, k, queries, got, want)
		}
	}
}
