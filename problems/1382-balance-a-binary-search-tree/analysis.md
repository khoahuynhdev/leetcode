# 1382. Balance a Binary Search Tree

[LeetCode Link](https://leetcode.com/problems/balance-a-binary-search-tree/)

Difficulty: Medium
Topics: Divide and Conquer, Greedy, Tree, Depth-First Search, Binary Search Tree, Binary Tree
Acceptance Rate: 84.9%

## Hints

### Hint 1

What property of a Binary Search Tree can you leverage to get all the values in sorted order? Once you have the values sorted, how might you build a balanced tree from them?

### Hint 2

An in-order traversal of a BST gives you the values in sorted order. Once you have a sorted array, think about how you would construct a balanced BST from it. Which element should be the root to ensure balance?

### Hint 3

The key insight is to use the middle element of a sorted array as the root of a balanced BST. This ensures equal (or near-equal) number of nodes in left and right subtrees. You can apply this recursively to build the entire tree.

## Approach

The solution follows a two-step approach:

**Step 1: In-order Traversal**
Perform an in-order traversal of the given BST to collect all node values in sorted order. Since in-order traversal of a BST visits nodes in ascending order, this gives us a sorted array of values.

**Step 2: Build Balanced BST from Sorted Array**
Use the sorted array to construct a balanced BST. The algorithm is:
1. Choose the middle element of the array (or subarray) as the root
2. Recursively build the left subtree from elements before the middle
3. Recursively build the right subtree from elements after the middle

This approach guarantees balance because at each step, we divide the remaining elements roughly equally between left and right subtrees. The height difference between any two subtrees will never exceed 1.

**Example walkthrough:**
For the skewed tree `[1,null,2,null,3,null,4]`:
1. In-order traversal gives: `[1, 2, 3, 4]`
2. Build balanced tree:
   - Middle element is 2 (index 1) → root
   - Left subarray `[1]` → middle is 1 → left child
   - Right subarray `[3, 4]` → middle is 3 → right child
     - Right subarray `[4]` → middle is 4 → right child of 3
3. Result: `[2,1,3,null,null,null,4]`

## Complexity Analysis

Time Complexity: O(n) where n is the number of nodes. We visit each node once during in-order traversal and once during tree construction.

Space Complexity: O(n) for storing the sorted array of node values. The recursion stack also uses O(log n) space for the balanced tree construction (or O(n) in worst case if we count the original skewed tree).

## Edge Cases

1. **Single node tree**: Tree is already balanced, but the algorithm handles it correctly by making it the root.
2. **Already balanced tree**: The algorithm will still reconstruct it, potentially in a different balanced configuration.
3. **Completely skewed tree**: The most common case this problem tests - ensures the algorithm properly balances a worst-case BST.
4. **Tree with duplicate values**: The problem constraints guarantee unique values (1 to 10^5), but the algorithm would handle duplicates correctly if they existed.
