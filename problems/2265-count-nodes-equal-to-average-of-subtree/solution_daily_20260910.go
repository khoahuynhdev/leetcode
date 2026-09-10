package main

// 2265. Count Nodes Equal to Average of Subtree
//
// Post-order DFS that returns the (sum, count) of each subtree. Averages cannot
// be merged across children of different sizes, but sums and counts can, so the
// recursion carries those raw ingredients upward and derives the average at each
// node. Since all values are non-negative, Go's integer division already floors,
// matching the rounding the problem asks for.
//
// Time: O(n), Space: O(h) for the recursion stack.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	answer := 0

	var dfs func(node *TreeNode) (sum, count int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		leftSum, leftCount := dfs(node.Left)
		rightSum, rightCount := dfs(node.Right)

		sum := node.Val + leftSum + rightSum
		count := 1 + leftCount + rightCount

		if node.Val == sum/count {
			answer++
		}

		return sum, count
	}

	dfs(root)

	return answer
}
