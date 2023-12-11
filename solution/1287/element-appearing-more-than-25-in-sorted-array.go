package solution

// naive implementation
// O(n) space
// O(n) time
func findSpecialInteger(arr []int) int {
	f := make(map[int]int)
	max := 1
	for _, v := range arr {
		f[v]++
		if f[max] < f[v] {
			max = v
		}
	}
	return max
}

// better solution since 25% equal 1/4 -> find a block that start and end with n/4
// O(1) space
// O(n) time
func findSpecialIntegerWithBlock(arr []int) int {
	size := len(arr) / 4
	for i := 0; i < len(arr); i++ {
		if arr[i] == arr[i+size] {
			return arr[i]
		}
	}
	return -1
}
