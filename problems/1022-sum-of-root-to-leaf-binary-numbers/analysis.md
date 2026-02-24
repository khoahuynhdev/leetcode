# 1022. Sum of Root To Leaf Binary Numbers

[LeetCode Link](https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/)

Difficulty: Easy
Topics: Tree, Depth-First Search, Binary Tree
Acceptance Rate: 74.1%

## Hints

### Hint 1

Think about how binary numbers work. As you traverse from root to leaf, each node you visit adds a new least-significant bit to the number being formed. What arithmetic operation converts appending a bit into a running numeric value?

### Hint 2

Consider using DFS to traverse every root-to-leaf path. As you go deeper, maintain a running integer that represents the binary number formed so far. When you move to a child node, shift the current value left by 1 and add the child's bit value — this is the same as `currentValue * 2 + node.val`.

### Hint 3

The base case is reaching a leaf node (a node with no children). At that point, the running value is the complete binary number for that path — add it to a total sum. For internal nodes, recursively pass the updated running value to both children and return the sum of results from the left and right subtrees.

## Approach

Use a recursive DFS traversal, carrying a `currentValue` parameter that accumulates the binary number as you descend.

1. Start at the root with `currentValue = 0`.
2. At each node, update: `currentValue = currentValue * 2 + node.Val`.
   - Multiplying by 2 is equivalent to a left bit-shift, making room for the new bit.
   - Adding `node.Val` (0 or 1) sets that new bit.
3. If the node is a leaf (both children are nil), return `currentValue` — this is the decimal value of one complete root-to-leaf binary number.
4. Otherwise, recurse into the left and right children and return the sum of their results.

**Example walkthrough** with `root = [1,0,1,0,1,0,1]`:

- Path `1→0→0`: currentValue goes `1 → 2 → 4`, which is binary `100` = 4
- Path `1→0→1`: currentValue goes `1 → 2 → 5`, which is binary `101` = 5
- Path `1→1→0`: currentValue goes `1 → 3 → 6`, which is binary `110` = 6
- Path `1→1→1`: currentValue goes `1 → 3 → 7`, which is binary `111` = 7
- Total = 4 + 5 + 6 + 7 = **22**

## Complexity Analysis

Time Complexity: O(n) — every node is visited exactly once.
Space Complexity: O(h) — where h is the height of the tree, due to the recursion stack. In the worst case (skewed tree) this is O(n); for a balanced tree it is O(log n).

## Edge Cases

- **Single node tree**: The root itself is a leaf. The answer is simply `root.Val` (0 or 1).
- **All zeros**: Every path produces the value 0, so the total sum is 0.
- **All ones**: Each path of depth d produces the value `2^d - 1`. Verify the sum is correct.
- **Left-skewed or right-skewed tree**: Only one root-to-leaf path exists. The recursion depth equals the number of nodes, which is the worst case for space.
