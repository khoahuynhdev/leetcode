package solution

// NOTE: Intuition: first I used recursion to calculate the result but soon hit the TLE
// -> for this we need to use a memoized function which basically stored the value after we computed 1
// -> avoid calculating too much value
// For this problem, there can be 2 ways
// -> top-down memoized
// -> bottom-up memoized
func tribonacci(n int) int {
	dict := make(map[int]int)
	var f func(int) int
	f = func(n int) int {
		if _, ok := dict[n]; ok {
			return dict[n]
		}
		if n == 0 {
			return 0
		}
		if n == 1 || n == 2 {
			return 1
		}

		lessss := f(n - 3)
		dict[n-3] = lessss
		lesss := f(n - 2)
		dict[n-2] = lesss
		less := f(n - 1)
		dict[n-1] = less
		return less + lesss + lessss
	}
	ans := f(n)
	return ans
}
