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

// using stack
func inorderTraversal(root *TreeNode) []int {
	// Follow up: Recursive solution is trivial, could you do it iteratively?
	s := make([]*TreeNode, 0)
	var ptr *TreeNode
	ptr = root
	ans := []int{}
	for ptr != nil {
		s = append([]*TreeNode{ptr}, s...)
		ptr = ptr.Left
	}
	for len(s) > 0 && ptr == nil {
		cur := s[0]
		ans = append(ans, cur.Val)
		s = s[1:]
		ptr = cur.Right
		for ptr != nil {
			s = append([]*TreeNode{ptr}, s...)
			ptr = ptr.Left
		}
	}
	return ans
}

// using recursion
func traverse(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, ans)
	*ans = append(*ans, root.Val)
	traverse(root.Right, ans)
}

func inorderTraversalRecursive(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	traverse(root, &ans)
	return ans
}
