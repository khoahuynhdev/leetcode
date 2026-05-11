package main

// Approach: For each integer in nums, extract its digits using repeated
// modulo/division (collecting them in reverse order), then append them to
// the result in correct left-to-right order by walking the temporary
// buffer backwards. We pre-allocate the result with a generous capacity
// since each input integer has at most 6 digits (nums[i] <= 10^5).

func separateDigits(nums []int) []int {
	result := make([]int, 0, len(nums)*6)
	var buf [6]int
	for _, n := range nums {
		idx := 0
		for n > 0 {
			buf[idx] = n % 10
			n /= 10
			idx++
		}
		for i := idx - 1; i >= 0; i-- {
			result = append(result, buf[i])
		}
	}
	return result
}
