# Binary Search Tree (BST) in Go

## What Makes a BST

A binary search tree is a binary tree where every node satisfies the BST invariant: all values in the left subtree are strictly less than the node's value, and all values in the right subtree are strictly greater. This ordering property is what gives BSTs their power — it lets you eliminate half the tree at each step during search, just like binary search on a sorted array.

```
        8
       / \
      3   10
     / \    \
    1   6    14
       / \   /
      4   7 13
```

The `TreeNode` definition used across LeetCode Go solutions:

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
```

## Traversals

Traversal order matters a lot for tree problems. Each order visits nodes in a specific sequence, and choosing the right one is often the key insight.

### In-Order (Left, Root, Right) — The BST Workhorse

In-order traversal visits a BST's nodes in sorted ascending order. This is the single most important property to remember for BST problems. Whenever a problem needs sorted values from a BST, in-order traversal is your tool.

```go
func inorder(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    inorder(node.Left, result)
    *result = append(*result, node.Val)
    inorder(node.Right, result)
}
```

Iterative version using an explicit stack (useful when you need to pause or control the traversal):

```go
func inorderIterative(root *TreeNode) []int {
    var result []int
    var stack []*TreeNode
    curr := root
    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, curr.Val)
        curr = curr.Right
    }
    return result
}
```

Used in: [1382. Balance a Binary Search Tree](../problems/1382-balance-a-binary-search-tree/solution.go) — collects sorted values, then rebuilds a balanced tree from the middle.

### Pre-Order (Root, Left, Right)

Visits the root before its children. Useful for serializing/copying a tree, since processing the root first naturally preserves structure. Also used when you need to pass information downward from parent to child (e.g., tracking the max value seen so far from root).

```go
func preorder(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    *result = append(*result, node.Val)
    preorder(node.Left, result)
    preorder(node.Right, result)
}
```

### Post-Order (Left, Right, Root)

Visits children before the root. This is the natural choice when you need to compute something bottom-up — the answer for a node depends on the answers for its children. Height calculation and balanced-tree checking both use post-order logic.

```go
func postorder(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    postorder(node.Left, result)
    postorder(node.Right, result)
    *result = append(*result, node.Val)
}
```

### Level-Order (BFS)

Visits nodes level by level using a queue. Essential for problems that care about depth or horizontal position.

```go
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    var result [][]int
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        levelSize := len(queue)
        var level []int
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            level = append(level, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        result = append(result, level)
    }
    return result
}
```

## Core BST Operations

### Search — O(h) where h is height

Follow the BST invariant: go left if the target is smaller, go right if larger. This is the fundamental operation that makes BSTs useful.

```go
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil || root.Val == val {
        return root
    }
    if val < root.Val {
        return searchBST(root.Left, val)
    }
    return searchBST(root.Right, val)
}
```

Related: [0700. Search in a Binary Search Tree](../problems/0700-search-in-a-binary-search-tree/solution.go)

### Insert — O(h)

Find the correct nil position by following the BST invariant, then place the new node there.

```go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    if val < root.Val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }
    return root
}
```

### Validate BST

A common mistake is checking only that `node.Left.Val < node.Val < node.Right.Val`. That's not enough — you need to ensure every node in the left subtree is less than the root, not just the immediate child. Pass down valid ranges instead.

```go
func isValidBST(root *TreeNode) bool {
    return validate(root, nil, nil)
}

func validate(node *TreeNode, min, max *int) bool {
    if node == nil {
        return true
    }
    if min != nil && node.Val <= *min {
        return false
    }
    if max != nil && node.Val >= *max {
        return false
    }
    return validate(node.Left, min, &node.Val) &&
        validate(node.Right, &node.Val, max)
}
```

An alternative approach: do an in-order traversal and check that the result is strictly increasing. Simpler to write, but uses O(n) extra space.

## Height and Balance

### Computing Height

Height is the number of edges on the longest path from a node to a leaf. A nil node has height -1 (or 0, depending on convention — LeetCode typically uses "depth" starting from 0 at the root).

```go
func height(node *TreeNode) int {
    if node == nil {
        return 0
    }
    left := height(node.Left)
    right := height(node.Right)
    if left > right {
        return left + 1
    }
    return right + 1
}
```

Related: [0104. Maximum Depth of Binary Tree](../problems/0104-maximum-depth-of-binary-tree/solution.go)

### Checking If Balanced

A tree is height-balanced if for every node, the heights of its left and right subtrees differ by at most 1. The naive approach computes height at every node (O(n log n)). The efficient approach computes height bottom-up and short-circuits as soon as imbalance is detected.

```go
func isBalanced(root *TreeNode) bool {
    return checkHeight(root) != -1
}

// Returns height if balanced, -1 if not
func checkHeight(node *TreeNode) int {
    if node == nil {
        return 0
    }
    left := checkHeight(node.Left)
    if left == -1 {
        return -1
    }
    right := checkHeight(node.Right)
    if right == -1 {
        return -1
    }
    if abs(left-right) > 1 {
        return -1
    }
    return max(left, right) + 1
}
```

Related: [0110. Balanced Binary Tree](../problems/0110-balanced-binary-tree/solution.go)

### Building a Balanced BST from Sorted Array

This is the core technique behind problem 1382. Given a sorted slice, always pick the middle element as root. This guarantees the tree is as balanced as possible because each subtree gets roughly half the elements.

```go
func buildBalancedBST(values []int, left, right int) *TreeNode {
    if left > right {
        return nil
    }
    mid := left + (right-left)/2
    node := &TreeNode{Val: values[mid]}
    node.Left = buildBalancedBST(values, left, mid-1)
    node.Right = buildBalancedBST(values, mid+1, right)
    return node
}
```

Related: [1382. Balance a Binary Search Tree](../problems/1382-balance-a-binary-search-tree/solution.go) — the full solution combines in-order traversal to get sorted values, then this function to rebuild.

## Finding Special Nodes

### Minimum and Maximum

In a BST, the minimum is the leftmost node and the maximum is the rightmost. Just keep going left (or right) until you hit nil.

```go
func findMin(node *TreeNode) *TreeNode {
    for node.Left != nil {
        node = node.Left
    }
    return node
}
```

### Lowest Common Ancestor (LCA) in BST

Unlike general binary trees where LCA requires checking both subtrees, BST structure gives you a shortcut: if both values are less than the current node, LCA is in the left subtree. If both are greater, it's in the right. Otherwise, the current node is the LCA.

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for root != nil {
        if p.Val < root.Val && q.Val < root.Val {
            root = root.Left
        } else if p.Val > root.Val && q.Val > root.Val {
            root = root.Right
        } else {
            return root
        }
    }
    return nil
}
```

### Kth Smallest Element

Do an in-order traversal and count. You can stop early once you've visited k nodes — no need to traverse the entire tree.

```go
func kthSmallest(root *TreeNode, k int) int {
    var stack []*TreeNode
    curr := root
    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        k--
        if k == 0 {
            return curr.Val
        }
        curr = curr.Right
    }
    return -1
}
```

## Tips and Tricks

### In-order traversal is your Swiss army knife for BST problems. If you're stuck on a BST problem, ask yourself "what would in-order traversal give me?" The sorted output often simplifies the problem dramatically. Problem 1382 is a perfect example — instead of complex tree rotations, just flatten to sorted array and rebuild.

### Know when to exploit BST ordering vs treating it as a generic tree. Some problems labeled "BST" don't actually need the ordering property (e.g., diameter, depth). Others are impossible without it (e.g., search, range sum). Read the problem carefully and ask: "does my solution change if I shuffle the values?" If no, you're not using the BST property, and you might be missing an optimization.

### Range queries are natural on BSTs. When a problem asks about nodes within a value range, use the BST ordering to prune entire subtrees. If `node.Val < low`, skip the left subtree entirely. If `node.Val > high`, skip the right subtree. This is the key insight for problems like [0938. Range Sum of BST](../problems/0938-range-sum-of-bts/solution.go).

### "Flatten and rebuild" is a legitimate strategy. For restructuring problems (balancing, merging two BSTs), it's often simplest to collect all values via in-order traversal into a sorted array, then build the desired structure from scratch. The time complexity is the same O(n), and the code is far simpler than in-place rotation approaches.

### Use pointer-to-pointer or return-value patterns for mutation. In Go, tree modification functions typically return the (possibly new) root, and the caller assigns it: `root.Left = insert(root.Left, val)`. This pattern handles the "node was nil and we created a new one" case cleanly.

### Height is a post-order computation. You need children's heights before you can compute a node's height. Whenever you see a problem that requires bottom-up information (height, subtree sum, is-balanced), think post-order DFS. Return the computed value up the call stack.

### For BST validation, pass constraints down — not just check neighbors. The valid range for each node is determined by all its ancestors, not just its parent. Pass `(min, max)` bounds through the recursion and tighten them as you descend.

### Iterative in-order traversal is worth memorizing. Many interview problems require modifying or stopping mid-traversal (kth smallest, BST iterator). The iterative stack-based version gives you explicit control over when to advance, which is hard to do with recursion.

## Complexity Reference

All operations are O(h) where h is the tree height. For a balanced BST, h = O(log n). For a skewed BST (essentially a linked list), h = O(n). This is why balancing matters — it's the difference between O(log n) and O(n) for every operation.

| Operation | Balanced BST | Skewed BST |
|---|---|---|
| Search | O(log n) | O(n) |
| Insert | O(log n) | O(n) |
| Delete | O(log n) | O(n) |
| Min/Max | O(log n) | O(n) |
| In-order traversal | O(n) | O(n) |

## Problems Log

| Problem | Category | Key Insight |
|---|---|---|
| [1382. Balance a Binary Search Tree](../problems/1382-balance-a-binary-search-tree/solution.go) | Flatten & Rebuild | In-order traversal gives sorted values; pick middle element as root recursively to build balanced BST. Avoids complex rotations entirely. |
| [0700. Search in a Binary Search Tree](../problems/0700-search-in-a-binary-search-tree/solution.go) | BST Search | Follow BST invariant: go left if smaller, right if larger. O(h) time. |
| [0938. Range Sum of BST](../problems/0938-range-sum-of-bts/solution.go) | Range Query | Use BST ordering to prune: skip left subtree if node < low, skip right if node > high. |
| [0104. Maximum Depth of Binary Tree](../problems/0104-maximum-depth-of-binary-tree/solution.go) | Height | Post-order: depth = 1 + max(left depth, right depth). |
| [0110. Balanced Binary Tree](../problems/0110-balanced-binary-tree/solution.go) | Balance Check | Single-pass post-order: return -1 sentinel to short-circuit on imbalance instead of computing height repeatedly. |
