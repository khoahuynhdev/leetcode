package main

import "testing"

func TestTotalWaviness(t *testing.T) {
	tests := []struct {
		name     string
		num1     int64
		num2     int64
		expected int64
	}{
		{"example 1: [120,130] has three peaks", 120, 130, 3},
		{"example 2: [198,202] mixes a peak and two valleys", 198, 202, 3},
		{"example 3: 4848 has one peak and one valley", 4848, 4848, 2},
		{"edge case: single-digit range has no neighbors", 1, 9, 0},
		{"edge case: all two-digit numbers are too short", 10, 99, 0},
		{"edge case: single number 121 is exactly one peak", 121, 121, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalWaviness(tt.num1, tt.num2); got != tt.expected {
				t.Errorf("totalWaviness(%d, %d) = %d, want %d", tt.num1, tt.num2, got, tt.expected)
			}
		})
	}
}

// wavinessBrute computes the waviness of a single number directly from its
// digits, used to cross-check the digit-DP solution.
func wavinessBrute(n int64) int64 {
	if n < 0 {
		return 0
	}
	var digits []int
	if n == 0 {
		digits = []int{0}
	} else {
		for n > 0 {
			digits = append(digits, int(n%10))
			n /= 10
		}
	}
	if len(digits) < 3 {
		return 0
	}
	// Peaks/valleys are symmetric under reversal, so scanning the (little-endian)
	// interior positions directly yields the same count.
	var w int64
	for i := 1; i < len(digits)-1; i++ {
		left, mid, right := digits[i-1], digits[i], digits[i+1]
		if (mid > left && mid > right) || (mid < left && mid < right) {
			w++
		}
	}
	return w
}

// sumWavinessBrute sums waviness over the inclusive range [a, b] by iterating.
func sumWavinessBrute(a, b int64) int64 {
	var total int64
	for x := a; x <= b; x++ {
		total += wavinessBrute(x)
	}
	return total
}

// TestTotalWavinessAgainstBruteForce cross-checks the digit DP against a direct
// summation over many ranges, including boundaries that exercise leading zeros
// and carries (crossing 99 -> 100, 999 -> 1000, 99999 -> 100000).
func TestTotalWavinessAgainstBruteForce(t *testing.T) {
	ranges := []struct {
		a, b int64
	}{
		{0, 1},
		{1, 200},
		{120, 130},
		{198, 202},
		{95, 105},
		{990, 1010},
		{1, 5000},
		{4848, 4848},
		{99990, 100010},
		{12321, 12321},
		{1, 12345},
		{50000, 60000},
	}

	for _, r := range ranges {
		want := sumWavinessBrute(r.a, r.b)
		if got := totalWaviness(r.a, r.b); got != want {
			t.Errorf("totalWaviness(%d, %d) = %d, brute force = %d", r.a, r.b, got, want)
		}
	}
}
