package main

// Approach: DFS over the original tree while keeping, for each letter, a stack
// of the ancestors on the current root-to-node path that carry that letter.
// When we enter node x, the top of stack[s[x]] (if any) is the closest ancestor
// with the same character, i.e. x's new parent. Because every new parent is an
// ancestor of x in the original tree, it appears earlier in the DFS preorder.
// So processing nodes in reverse preorder and adding each node's size to its
// new parent's size yields the final subtree sizes.
//
// Time: O(n), Space: O(n).
func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		children[parent[i]] = append(children[parent[i]], i)
	}

	newParent := make([]int, n)
	copy(newParent, parent)

	var last [26][]int
	order := make([]int, 0, n)

	var dfs func(x int)
	dfs = func(x int) {
		order = append(order, x)
		c := s[x] - 'a'
		if k := len(last[c]); k > 0 {
			newParent[x] = last[c][k-1]
		}
		last[c] = append(last[c], x)
		for _, ch := range children[x] {
			dfs(ch)
		}
		last[c] = last[c][:len(last[c])-1]
	}
	dfs(0)

	answer := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		x := order[i]
		answer[x]++
		if p := newParent[x]; p >= 0 {
			answer[p] += answer[x]
		}
	}
	return answer
}
