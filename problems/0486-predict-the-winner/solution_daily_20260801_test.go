package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "example 1: [1,5,2] player 1 always hands over the 5",
			nums:     []int{1, 5, 2},
			expected: false,
		},
		{
			name:     "example 2: [1,5,233,7] player 1 sacrifices the 1 to grab 233",
			nums:     []int{1, 5, 233, 7},
			expected: true,
		},
		{
			name:     "edge case: single element, player 2 never moves",
			nums:     []int{5},
			expected: true,
		},
		{
			name:     "edge case: tie counts as a win for player 1",
			nums:     []int{1, 1},
			expected: true,
		},
		{
			name:     "edge case: all zeros ties at 0-0",
			nums:     []int{0, 0, 0},
			expected: true,
		},
		{
			name:     "edge case: empty input, no moves at all",
			nums:     []int{},
			expected: true,
		},
		{
			name:     "odd length: [5,1,2] leading big value is safe",
			nums:     []int{5, 1, 2},
			expected: true,
		},
		{
			name:     "odd length: [2,4,55,6,8] the 55 is unreachable for player 1",
			nums:     []int{2, 4, 55, 6, 8},
			expected: false,
		},
		{
			name:     "even length is always winnable for player 1",
			nums:     []int{1, 2, 3, 4},
			expected: true,
		},
		{
			name:     "edge case: maximum constraint values, no overflow",
			nums:     []int{10000000, 1, 10000000, 1},
			expected: true,
		},
		{
			name:     "edge case: 20 elements, largest allowed input",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PredictTheWinner(tt.nums)
			if result != tt.expected {
				t.Errorf("PredictTheWinner(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

// bruteForceDiff exhaustively explores the game tree and returns the best
// achievable (mover - opponent) score difference. It is exponential, so it is
// only used on short inputs to cross-check the DP.
func bruteForceDiff(nums []int, i, j int) int {
	if i > j {
		return 0
	}
	takeLeft := nums[i] - bruteForceDiff(nums, i+1, j)
	takeRight := nums[j] - bruteForceDiff(nums, i, j-1)
	if takeLeft > takeRight {
		return takeLeft
	}
	return takeRight
}

func TestSolutionAgainstBruteForce(t *testing.T) {
	inputs := [][]int{
		{1, 5, 2},
		{1, 5, 233, 7},
		{5, 1, 2},
		{2, 4, 55, 6, 8},
		{0, 0, 0, 0},
		{3, 9, 1, 2},
		{7, 4, 4, 7, 1},
		{1, 2, 3, 4, 5, 6},
		{10, 1, 1, 10, 1, 1, 10},
		{9},
	}

	for _, nums := range inputs {
		nums := nums
		t.Run("brute force parity", func(t *testing.T) {
			want := bruteForceDiff(nums, 0, len(nums)-1) >= 0
			if got := PredictTheWinner(nums); got != want {
				t.Errorf("PredictTheWinner(%v) = %v, brute force says %v", nums, got, want)
			}
		})
	}
}

func TestSolutionDoesNotMutateInput(t *testing.T) {
	nums := []int{1, 5, 233, 7}
	original := append([]int(nil), nums...)

	PredictTheWinner(nums)

	for i := range nums {
		if nums[i] != original[i] {
			t.Fatalf("input mutated: got %v, want %v", nums, original)
		}
	}
}
