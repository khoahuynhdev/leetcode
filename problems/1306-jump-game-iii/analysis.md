# 1306. Jump Game III

[LeetCode Link](https://leetcode.com/problems/jump-game-iii/)

Difficulty: Medium
Topics: Array, Depth-First Search, Breadth-First Search
Acceptance Rate: 67.5%

## Hints

### Hint 1

Think of the array as a graph. Each index `i` is a node, and from it you have at most two outgoing edges: one to `i + arr[i]` and one to `i - arr[i]` (when those land inside the array). The question becomes a simple reachability problem: starting at `start`, can you reach any node whose value is `0`? When you see "reach any node from a starting node", the toolbox is graph traversal — BFS or DFS.

### Hint 2

A naive recursive exploration can revisit the same index endlessly (jumping back and forth between two indices). You need a way to mark indices you have already explored so each one is processed at most once. A boolean `visited` array (or repurposing the input by marking values negative) is enough. Once you bound work to "each index at most once", the algorithm is linear in the array length.

### Hint 3

The actual jump values do not need to be precomputed or sorted; you simply follow them. DFS with a visited set is the most concise implementation: at each step check if the current cell is `0` (success), otherwise mark it visited and recurse into both neighbors. Return `true` as soon as any branch succeeds. Bounds checking on `i + arr[i]` and `i - arr[i]` keeps you from jumping outside the array. Because every index is touched at most once, the search terminates even when the answer is `false`.

## Approach

Model the array as an implicit graph and run a depth-first search from `start`:

1. Create a `visited` array of size `n`.
2. Define a recursive helper `dfs(i)`:
   - If `i` is out of bounds or already visited, return `false`.
   - If `arr[i] == 0`, return `true` — we have reached a zero cell.
   - Mark `i` as visited.
   - Recurse into `dfs(i + arr[i])` and `dfs(i - arr[i])`; return `true` if either succeeds.
3. Return the result of `dfs(start)`.

Why it works: any reachable zero must lie on some sequence of legal jumps from `start`. DFS explores all such sequences. The `visited` array prevents infinite loops (e.g. two indices whose jumps point at each other) and ensures every index is examined at most once, which keeps the algorithm linear.

Walk through Example 1, `arr = [4,2,3,0,3,1,2]`, `start = 5`:
- `dfs(5)`: `arr[5] = 1`, mark 5. Try `dfs(6)` then `dfs(4)`.
- `dfs(6)`: `arr[6] = 2`, mark 6. Try `dfs(8)` (out of bounds, false) then `dfs(4)`.
- `dfs(4)`: `arr[4] = 3`, mark 4. Try `dfs(7)` (out of bounds, false) then `dfs(1)`.
- `dfs(1)`: `arr[1] = 2`, mark 1. Try `dfs(3)`.
- `dfs(3)`: `arr[3] = 0` → return `true`.

The chain unwinds and the original call returns `true`.

A BFS using a queue is equivalent — pick whichever you find clearer. DFS via recursion is slightly more compact in Go for this problem.

## Complexity Analysis

Time Complexity: O(n) — each index is visited at most once, and each visit does O(1) work.
Space Complexity: O(n) — the `visited` array plus the recursion stack, both bounded by the array length.

## Edge Cases

- **`arr[start] == 0`**: the answer is immediately `true`; the base case in DFS handles this on the first call.
- **No zero present in the array**: traversal will exhaust all reachable indices and return `false`. The `visited` array guarantees termination.
- **Single-element array**: by the constraint `0 <= arr[i] < arr.length`, the lone element must be `0`, so the answer is `true`.
- **Disconnected zero**: a zero exists but lies in a component the start cannot reach. DFS will visit only the start's component, never touch the zero, and correctly return `false` (Example 3 is this case).
- **Jumps that land out of bounds**: `i + arr[i]` or `i - arr[i]` may fall outside `[0, n)`. The bounds check at the top of `dfs` rejects them safely instead of panicking on an index-out-of-range.
