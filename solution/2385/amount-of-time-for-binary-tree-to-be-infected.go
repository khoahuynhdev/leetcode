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

// NOTE: my false assumption
// The number of nodes in the tree is in the range [1, 1e5] -> not necessary the maximum size of the array
// if using heap solution -> large dataset will cause MLE
func buildGraph(node *TreeNode, g map[int][]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		g[node.Val] = append(g[node.Val], node.Left.Val)
		g[node.Left.Val] = append(g[node.Left.Val], node.Val)
		buildGraph(node.Left, g)
	}
	if node.Right != nil {
		g[node.Val] = append(g[node.Val], node.Right.Val)
		g[node.Right.Val] = append(g[node.Right.Val], node.Val)
		buildGraph(node.Right, g)
	}
}

func amountOfTime(root *TreeNode, start int) int {
	time := -1
	if root.Left == nil && root.Right == nil {
		return 0
	}
	var infected map[int]bool = map[int]bool{}
	var g map[int][]int = map[int][]int{}
	buildGraph(root, g)
	var q []int = []int{start}
	for len(q) > 0 {
		size := len(q)
		time++
		for i := 0; i < size; i++ {
			node := q[i]
			infected[node] = true
			// note: check for nil before adding to q
			for _, adj := range g[node] {
				if !infected[adj] {
					q = append(q, adj)
				}
			}
		}
		q = q[size:]
		// fmt.Println(q, g,infected)
	}
	return time
}
