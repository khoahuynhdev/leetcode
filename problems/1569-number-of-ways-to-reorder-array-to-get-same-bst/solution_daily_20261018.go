package main

// Approach: divide and conquer + combinatorics.
// The first element is always the root. The rest split into a "left" list
// (smaller values) and a "right" list (larger values), each keeping its
// relative order. Any interleaving of those two lists preserves the BST, so
//
//	ways(nums) = C(len(left)+len(right), len(left)) * ways(left) * ways(right)
//
// Binomial coefficients come from Pascal's triangle (mod 1e9+7). The answer
// subtracts 1 to exclude the original ordering.
// Time: O(n^2), Space: O(n^2) for the Pascal table.

const mod = 1_000_000_007

func numOfWays(nums []int) int {
	n := len(nums)

	comb := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		comb[i] = make([]int, i+1)
		comb[i][0], comb[i][i] = 1, 1
		for j := 1; j < i; j++ {
			comb[i][j] = (comb[i-1][j-1] + comb[i-1][j]) % mod
		}
	}

	var ways func(arr []int) int
	ways = func(arr []int) int {
		if len(arr) <= 2 {
			return 1
		}
		root := arr[0]
		var left, right []int
		for _, v := range arr[1:] {
			if v < root {
				left = append(left, v)
			} else {
				right = append(right, v)
			}
		}
		res := comb[len(arr)-1][len(left)]
		res = res * ways(left) % mod
		res = res * ways(right) % mod
		return res
	}

	return (ways(nums) - 1 + mod) % mod
}
