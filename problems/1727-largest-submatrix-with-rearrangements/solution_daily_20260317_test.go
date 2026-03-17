package solution

import "testing"

func TestLargestSubmatrixDaily(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected int
	}{
		{
			name:     "example 1: 3x3 with rearrangement giving 4",
			matrix:   [][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}},
			expected: 4,
		},
		{
			name:     "example 2: single row with three 1s",
			matrix:   [][]int{{1, 0, 1, 0, 1}},
			expected: 3,
		},
		{
			name:     "example 3: 2x3 best area is 2",
			matrix:   [][]int{{1, 1, 0}, {1, 0, 1}},
			expected: 2,
		},
		{
			name:     "edge case: all zeros",
			matrix:   [][]int{{0, 0}, {0, 0}},
			expected: 0,
		},
		{
			name:     "edge case: all ones",
			matrix:   [][]int{{1, 1, 1}, {1, 1, 1}},
			expected: 6,
		},
		{
			name:     "edge case: single element 1",
			matrix:   [][]int{{1}},
			expected: 1,
		},
		{
			name:     "edge case: single element 0",
			matrix:   [][]int{{0}},
			expected: 0,
		},
		{
			name:     "edge case: single column with consecutive ones",
			matrix:   [][]int{{1}, {1}, {1}, {0}, {1}},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Deep copy matrix since the function modifies it in place.
			matrix := make([][]int, len(tt.matrix))
			for i := range tt.matrix {
				matrix[i] = make([]int, len(tt.matrix[i]))
				copy(matrix[i], tt.matrix[i])
			}
			result := largestSubmatrixDaily(matrix)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
