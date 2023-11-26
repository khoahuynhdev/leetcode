package solution

// O(n) time
// O(n) space
// grouping
//
// my naive solution match but not pass the last testcase
//
//	func findDiagonalOrder(nums [][]int) []int {
//		size := len(nums)
//	    if size == 1 {
//	        return nums[0]
//	    }
//		ans := []int{}
//		for i := 1; i <= (2*size)-1; i++ {
//			m := int(math.Min(float64(i-1), float64(size)))
//			n := 0
//			if i > size {
//				n = i % size
//				m = size - 1
//			}
//			for j := size - int(math.Abs(float64(size-i))); j > 0; j-- {
//	            //fmt.Printf("i:%d;j:%d;m:%d;n:%d\n", i, j,m,n)
//				if n >= len(nums[m]) {
//	                n++
//				    m--
//					continue
//				}
//				ans = append(ans, nums[m][n])
//				n++
//				m--
//			}
//		}
//		return ans
//	}
//
// every element where sum(i+j) equal belong to a diagonal -> use a hashmap to store them
// To collect the cells on each diagonal in the correct order,
// we will iterate through each row from left to right starting with the bottom row.
// The reason we choose the bottom-up, left-to-right order is that
// the diagonals move upward and to the right, so by iterating to the upper right,
// we will visit the squares in the correct order.
// =>>> the iterate will depend on the direction of the diagonal
func findDiagonalOrder(nums [][]int) []int {
	ans := []int{}
	groups := make(map[int][]int)
	n := 0
	for row := len(nums) - 1; row >= 0; row-- {
		for col := 0; col < len(nums[row]); col++ {
			diagonal := row + col
			groups[diagonal] = append(groups[diagonal], nums[row][col])
			n++
		}
	}

	curr := 0
	for len(groups[curr]) > 0 {
		for _, num := range groups[curr] {
			ans = append(ans, num)
		}
		curr++
	}
	return ans
}
