# 3619. Count Islands With Total Value Divisible by K

[LeetCode Link](https://leetcode.com/problems/count-islands-with-total-value-divisible-by-k/)

Difficulty: Medium
Topics: Array, Depth-First Search, Breadth-First Search, Union-Find, Matrix
Acceptance Rate: 59.9%

## Hints

### Hint 1

This is a twist on the classic "Number of Islands" problem (LeetCode 200). What technique do you use there to find every connected group of land cells?

### Hint 2

Use a flood fill (DFS or BFS) starting from each unvisited land cell. The only change from the classic problem is that you need to carry extra information out of each flood fill.

### Hint 3

While flood-filling an island, add up the values of every cell you visit. When the fill finishes, you have the island's total. Check `total % k == 0` and count it. Mark cells visited (e.g. set them to `0`) so no cell is counted twice.

## Approach

1. Scan every cell of the grid in row-major order.
2. When you find a positive cell, it is the first cell you have seen of a new island. Start a flood fill from it:
   - Push it onto a stack, add its value to `sum`, and set it to `0` to mark it visited.
   - Pop cells one at a time; for each of the 4 neighbors that is in bounds and still positive, add its value to `sum`, zero it, and push it.
3. When the stack is empty, the whole island has been explored and `sum` is its total value. If `sum % k == 0`, increment the answer.
4. Keep scanning. Cells that were zeroed are skipped, so each island is processed exactly once.

Why it works: a flood fill reaches exactly the cells that are 4-directionally connected to the start cell through positive cells, which is the definition of an island. Zeroing a cell when it is pushed (not when it is popped) means it can never be pushed twice, so each value is added exactly once.

An explicit stack is used instead of recursion. A single island can hold up to 10^5 cells, and an iterative fill avoids deep recursion. Go's growable goroutine stacks would handle recursion here too, but iterative is the safer habit and easy to carry over to other languages.

Example 2: `[[3,0,3,0],[0,3,0,3],[3,0,3,0]]`, `k = 3`. Each `3` touches the others only diagonally, so there are 6 single-cell islands, each with total 3, and all 6 are divisible by 3.

Alternative: Union-Find works too. Union adjacent land cells and keep a per-root running sum, then count the roots whose sum is divisible by `k`. It is more code for the same asymptotic cost.

## Complexity Analysis

Time Complexity: O(m * n). Each cell is visited once by the scan and pushed or popped at most once.
Space Complexity: O(m * n) in the worst case for the stack (one island that covers the whole grid). Marking visited in place needs no extra visited array.

## Edge Cases

- **No land at all**: the answer is 0, since no flood fill ever starts.
- **Diagonal neighbors**: diagonal cells are *not* connected. Example 2 relies on this, and using 8 directions by mistake gives the wrong answer.
- **Overflow**: up to 10^5 cells × 10^6 per cell = 10^11, which overflows 32-bit integers. Go's `int` is 64-bit on LeetCode, but in languages like Java or C++ use `long`, or keep `sum % k` as you go.
- **`k = 1`**: every island counts.
- **One island covering the whole grid**: this is the deepest possible fill. It is where recursive DFS is most likely to overflow the stack.
- **Mutating input**: zeroing cells destroys the grid. If the caller needs it afterward, use a separate `visited` array.
