# 0110. Balanced Binary Tree

[LeetCode Link](https://leetcode.com/problems/balanced-binary-tree/)

Difficulty: Easy
Topics: Tree, Depth-First Search, Binary Tree
Acceptance Rate: 57.1%

## Hints

### Hint 1

A height-balanced binary tree means that for every node, the heights of its left and right subtrees differ by at most 1. Think about how you would compute the height of a tree and what information you need to check at each node.

### Hint 2

You can use a depth-first search (DFS) approach. Consider calculating the height of each subtree while simultaneously checking the balance condition. What should you do if you discover an imbalance?

### Hint 3

The key insight is to combine height calculation with balance checking in a single pass. Return a special value (like -1) when a subtree is unbalanced, which allows you to short-circuit and avoid unnecessary computation once an imbalance is detected.

## Approach

The solution uses a bottom-up DFS approach that calculates height and checks balance simultaneously:

1. **Base Case**: An empty tree (nil node) is balanced with height 0.

2. **Recursive Case**: For each node:
   - Recursively calculate the height of the left subtree
   - Recursively calculate the height of the right subtree
   - If either subtree is unbalanced (returns -1), propagate the unbalanced state upward
   - Check if the current node is balanced by comparing left and right heights
   - If the difference is greater than 1, return -1 (unbalanced)
   - Otherwise, return the height of the current node (max of left and right heights + 1)

3. **Final Check**: If the helper function returns -1, the tree is not balanced; otherwise, it is balanced.

This approach is efficient because it:
- Visits each node exactly once
- Checks the balance condition while computing heights
- Short-circuits early when an imbalance is detected

**Example Walkthrough** (for `[3,9,20,null,null,15,7]`):
- Node 9: height = 1 (leaf node, balanced)
- Node 15: height = 1 (leaf node, balanced)
- Node 7: height = 1 (leaf node, balanced)
- Node 20: left height = 1, right height = 1, difference = 0 (balanced), height = 2
- Node 3: left height = 1, right height = 2, difference = 1 (balanced), height = 3

## Complexity Analysis

Time Complexity: O(n), where n is the number of nodes in the tree. We visit each node exactly once.

Space Complexity: O(h), where h is the height of the tree. This space is used by the recursion call stack. In the worst case (skewed tree), h = n, making it O(n). In the best case (balanced tree), h = log(n).

## Edge Cases

- **Empty tree**: A nil root should return true (an empty tree is balanced by definition).
- **Single node**: A tree with only one node is always balanced.
- **Completely skewed tree**: A tree that looks like a linked list (all nodes in one direction) is technically balanced if there's only one branch.
- **Subtree imbalance**: The imbalance might occur deep in the tree, not just at the root level.
- **Multiple imbalances**: Once one imbalance is found, we should stop checking and return false immediately.
