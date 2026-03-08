package solution

import "testing"

func TestFindDifferentBinaryString(t *testing.T) {
	tests := []struct {
		name     string
		nums     []string
		notInSet bool // true means we validate result is not in nums rather than exact match
		valid    map[string]bool
	}{
		{
			name: "example 1: two binary strings",
			nums: []string{"01", "10"},
			valid: map[string]bool{
				"00": true,
				"11": true,
			},
		},
		{
			name: "example 2: consecutive binary strings",
			nums: []string{"00", "01"},
			valid: map[string]bool{
				"10": true,
				"11": true,
			},
		},
		{
			name: "example 3: three binary strings",
			nums: []string{"111", "011", "001"},
			valid: map[string]bool{
				"000": true,
				"010": true,
				"100": true,
				"101": true,
				"110": true,
			},
		},
		{
			name: "edge case: single string 0",
			nums: []string{"0"},
			valid: map[string]bool{
				"1": true,
			},
		},
		{
			name: "edge case: single string 1",
			nums: []string{"1"},
			valid: map[string]bool{
				"0": true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findDifferentBinaryStringDaily(tt.nums)
			if len(result) != len(tt.nums) {
				t.Errorf("got length %d, want length %d", len(result), len(tt.nums))
				return
			}
			if !tt.valid[result] {
				// Check that the result is at least not in nums
				for _, s := range tt.nums {
					if result == s {
						t.Errorf("result %q is in nums, expected a string not in nums", result)
						return
					}
				}
			}
		})
	}
}
