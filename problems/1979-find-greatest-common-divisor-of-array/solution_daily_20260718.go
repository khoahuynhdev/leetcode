package main

// Approach: Scan the array once to find the smallest and largest values, then
// compute their greatest common divisor using the Euclidean algorithm
// (gcd(a, b) = gcd(b, a % b) until b == 0). This avoids sorting and runs in
// O(n) time with O(1) extra space.

func findGCD(nums []int) int {
	min, max := nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return gcd(max, min)
}

// gcd returns the greatest common divisor of a and b via the Euclidean algorithm.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
