package solution

// https://leetcode.com/problems/maximum-product-difference-between-two-pairs/description/?envType=daily-question&envId=2023-12-18

// Intuition: this problem is rather easy, can be solved with O(n^2)
// or Sort with O(nlogn)
// My approach is a big constant -> O(1) time
func maxProductDifference(nums []int) int {
	minC, maxC := 0, 0
	minProd, maxProd := 1, 1
	vals := make([]int, 10001)
	for _, v := range nums {
		vals[v] += 1
	}
	for i, j := 0, len(vals)-1; i < len(vals) && j >= 0; i, j = i+1, j-1 {
		if vals[i] > 0 {
			for vals[i] > 0 && minC < 2 {
				minProd *= i
				minC++
				vals[i]--
			}
		}
		if vals[j] > 0 {
			for vals[j] > 0 && maxC < 2 {
				maxProd *= j
				maxC++
				vals[j]--
			}
		}
		if minC >= 2 && maxC >= 2 {
			break
		}
	}
	return maxProd - minProd
}
