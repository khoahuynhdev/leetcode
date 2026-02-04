package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsPseudoPalindrom(arr []int) bool {
	hasOdd := false
	for i := 1; i < 10; i++ {
		if arr[i]%2 == 1 {
			if hasOdd {
				return false
			} else {
				hasOdd = true
			}
		}
	}
	return true
}

func traverse(n *TreeNode, arr []int) int {
	if n == nil {
		return 0
	}
	if n.Left == nil && n.Right == nil {
		arr[n.Val]++
		isPalindrome := IsPseudoPalindrom(arr)
		arr[n.Val]--
		if isPalindrome {
			return 1
		} else {
			return 0
		}
	}
	arr[n.Val]++
	l, r := traverse(n.Left, arr), traverse(n.Right, arr)
	arr[n.Val]--
	return l + r
}

func pseudoPalindromicPaths(root *TreeNode) int {
	vals := make([]int, 10)
	ans := traverse(root, vals)
	return ans
}
