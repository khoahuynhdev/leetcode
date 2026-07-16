package main

import "testing"

func TestSumGcd(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "example 1: [2,6,4] -> sorted [2,2,6], gcd(2,6)=2",
			nums:     []int{2, 6, 4},
			expected: 2,
		},
		{
			name:     "example 2: [3,6,2,8] -> sorted [2,3,6,8], gcd(2,8)+gcd(3,6)=5",
			nums:     []int{3, 6, 2, 8},
			expected: 5,
		},
		{
			name:     "edge case: single element has no pairs",
			nums:     []int{7},
			expected: 0,
		},
		{
			name:     "edge case: two identical elements, gcd(5,5)=5",
			nums:     []int{5, 5},
			expected: 5,
		},
		{
			name:     "edge case: odd length middle ignored",
			nums:     []int{5, 5, 5, 5, 5},
			expected: 10, // sorted [5,5,5,5,5]: gcd(5,5)+gcd(5,5)=10, middle skipped
		},
		{
			name:     "edge case: strictly increasing keeps values as prefixGcd",
			nums:     []int{1, 2, 3, 4},
			expected: 2, // prefixGcd [1,2,3,4]: gcd(1,4)+gcd(2,3)=1+1=2
		},
		{
			name:     "edge case: large values, no overflow",
			nums:     []int{1000000000, 1000000000},
			expected: 1000000000, // gcd(1e9,1e9)=1e9
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumGcd(tt.nums)
			if result != tt.expected {
				t.Errorf("sumGcd(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
