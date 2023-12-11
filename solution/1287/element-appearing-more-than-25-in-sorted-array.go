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

// O(1) space
// O(logn) time
func findSpecialIntegerPromax(arr []int) int {
	size := len(arr) / 4
	n := len(arr)
	if n == 1 {
		return 1
	}
	candidates := []int{arr[n/4], arr[n/2], arr[3*n/4]}
	for _, candidate := range candidates {
		left := IntLowerBound(arr, candidate)
		right := IntUpperBound(arr, candidate) - 1
		if right-left+1 > size {
			return candidate
		}
	}
	return -1
}

func IntUpperBound(arr []int, target int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func IntLowerBound(arr []int, target int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
