package solution

// intuition: you can do this
// https://leetcode.com/problems/find-the-highest-altitude/?envType=study-plan-v2&envId=leetcode-75
func largestAltitude(gain []int) int {
	max := 0
	gained := make([]int, len(gain)+1)
	gained[0] = 0
	for i := 0; i < len(gain); i++ {
		gained[i+1] = gained[i] + gain[i]
		if max < gained[i+1] {
			max = gained[i+1]
		}
	}
	return max
}
