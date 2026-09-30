package main

// Approach: track the current nesting depth while scanning. A '(' and its
// matching ')' share the same depth, so assigning each parenthesis to group
// depth%2 keeps both groups valid and alternates levels between them.
// Each group gets at most ceil(D/2) levels, which is optimal.
// Time: O(n), Space: O(1) extra (excluding output).
func maxDepthAfterSplit(seq string) []int {
	ans := make([]int, len(seq))
	depth := 0
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			depth++
			ans[i] = depth % 2
		} else {
			ans[i] = depth % 2
			depth--
		}
	}
	return ans
}
