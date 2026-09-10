# 2265. Count Nodes Equal to Average of Subtree

[LeetCode Link](https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/)

Difficulty: Medium
Topics: Tree, Depth-First Search, Binary Tree
Acceptance Rate: 87.4%

## Hints

### Hint 1

The question asked at every node ("what is the average of my subtree?") depends only on information that lives *below* that node. Whenever a node's answer is a function of its children's answers, think bottom-up recursion — a post-order DFS.

### Hint 2

An average is not something you can combine directly: knowing the left subtree's average and the right subtree's average is not enough to compute the parent's average, because the two subtrees may have different sizes. Ask yourself what pair of numbers *is* composable across children.

### Hint 3

Have the recursion return **two** values per subtree: its sum and its node count. Both compose trivially — `sum = node.Val + leftSum + rightSum` and `count = 1 + leftCount + rightCount`. Once you have them at a node, check `node.Val == sum / count` (Go's integer division on non-negative values already floors, which is exactly the rounding the problem wants) and increment a counter. One traversal computes every node's answer.

## Approach

This is the classic "return an aggregate from DFS" pattern, and the whole trick is choosing an aggregate that merges cleanly.

Define a recursive helper `dfs(node) (sum, count int)`:

1. **Base case:** a `nil` node contributes nothing, so return `(0, 0)`. This makes nodes with one missing child fall out of the general case with no special handling.
2. **Recurse:** get `(leftSum, leftCount)` and `(rightSum, rightCount)` from the children.
3. **Combine:** `sum = node.Val + leftSum + rightSum`, `count = 1 + leftCount + rightCount`.
4. **Check the node itself:** if `node.Val == sum/count`, bump a shared counter. `count` is at least 1 here (the node itself), so the division is always safe.
5. **Return** `(sum, count)` so the parent can build its own aggregate.

The answer is the counter's value after the traversal. A closure over a captured `answer` variable keeps the helper's signature focused on the aggregate rather than threading the count through the return values as well.

Why the floor is free: the constraints guarantee `0 <= Node.val <= 1000`, so every subtree sum is non-negative, and Go's `/` on non-negative integers truncates toward zero — which equals flooring. If values could be negative you would have to floor explicitly, since Go truncates toward zero rather than toward negative infinity.

Walking Example 1, `root = [4,8,5,0,1,null,6]`:

- Node `0` → `(0, 1)`, average `0` → matches, count it.
- Node `1` → `(1, 1)`, average `1` → matches.
- Node `8` → sum `8+0+1 = 9`, count `3`, average `9/3 = 3` ≠ `8` → no.
- Node `6` → `(6, 1)`, average `6` → matches.
- Node `5` → sum `5+6 = 11`, count `2`, average `11/2 = 5` (floored) → matches. This is the case that makes the flooring rule visible.
- Node `4` (root) → sum `24`, count `6`, average `4` → matches.

Total: 5, as expected.

Overflow is a non-issue: at most 1000 nodes with values up to 1000 caps the sum at 1,000,000, far inside `int`.

## Complexity Analysis

Time Complexity: O(n) — each node is visited exactly once and does O(1) work.
Space Complexity: O(h) for the recursion stack, where `h` is the tree height. That is O(log n) for a balanced tree and O(n) in the worst case of a degenerate, list-like tree (with n up to 1000, recursion depth is never a concern here).

## Edge Cases

- **Single node tree** (`[1]`): the subtree average is the node's own value, so the answer is always 1. A good sanity check that the base case returns `(0, 0)` rather than something that skews the sum or count.
- **`nil` children / nodes with only one child**: the `(0, 0)` base case must contribute nothing to either the sum or the count. Returning a count of 1 for `nil` would silently deflate every average above it.
- **Nodes with value 0**: `0 == 0/1` matches, so zero-valued leaves always count. Don't let a truthiness-style check or a "skip empty values" instinct drop them.
- **Flooring matters**: subtrees like `[5, null, 6]` (sum 11, count 2) only match because `11/2` floors to `5`. Using floating-point division and comparing to `node.Val` would reject this node.
- **Skewed trees**: a 1000-node chain gives recursion depth 1000, which is fine in Go; worth noting only if you were asked to handle much deeper trees, where an explicit stack or Morris-style traversal would be needed.

Honest difficulty note: the traversal is short, but the reason this is a Medium rather than an Easy is the "averages don't compose, sums and counts do" insight. If you internalize that — return the raw ingredients from DFS, not the derived value — a whole family of tree problems opens up.
