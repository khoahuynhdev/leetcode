package solution

// intuition
// mininum number of operation -> using greedy
// Time: O(n)
// Space: O(n)
// check for occurences < 2
// check for occurences % 3 == 0
// check for occurences % 3 == 1
// check for occurences % 3 == 2
func minOperations(nums []int) int {
	ans := 0
	f := map[int]int{}
	for _, v := range nums {
		f[v]++
	}
	for _, v := range f {
		if v < 2 {
			return -1
		}
		if v%3 == 0 {
			ans += v / 3
		}
		if v%3 == 2 {
			ans += (v / 3) + 1
		}
		if v%3 == 1 {
			ans += ((v - 2) / 3) + 2
		}
	}
	return ans
}
