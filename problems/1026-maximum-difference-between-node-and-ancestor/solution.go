package solution

// https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/submissions/1143379157/?envType=daily-question&envId=2024-01-11
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

// debug for too many state
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func dfsm(root *TreeNode, m int, ma int) (int, int, *int) {
	if root == nil {
		hold := -1
		return m, ma, &hold
	}
	minL, maxL, pl := dfsm(root.Left, root.Val, root.Val)
	minR, maxR, pr := dfsm(root.Right, root.Val, root.Val)
	c := min(root.Val, min(minL, minR))
	d := max(root.Val, max(maxL, maxR))

	newMax := max(max(*pl, *pr), max(abs(min(minL, minR), root.Val), abs(max(maxR, maxL), root.Val)))
	// fmt.Println(newMax, min(minL,minR),max(maxL,maxR), root.Val)
	return min(c, m), max(ma, d), &newMax
}

func maxAncestorDiff(root *TreeNode) int {
	_, _, a := dfsm(root, root.Val, 0)
	return *a
}
