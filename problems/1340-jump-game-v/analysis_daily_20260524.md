# 1340. Jump Game V

[LeetCode Link](https://leetcode.com/problems/jump-game-v/)

Difficulty: Hard
Topics: Array, Dynamic Programming, Sorting
Acceptance Rate: 68.5%

## Hints

### Hint 1

Notice that you may *start* at any index, and you want the longest chain of jumps. Whenever a problem asks "what is the best result achievable from a starting point, where the result depends on the best result of the next points you can reach," that is a strong signal for **dynamic programming**. Try to define a value for each index that captures "the most indices I can visit if I begin here."

### Hint 2

Every legal jump goes from a *higher* bar to a *strictly lower* one (`arr[i] > arr[j]`). That means the jump relation is **acyclic** — you can never return to where you came from, and you can never reach an index whose value is `>= arr[i]`. So `dp[i]` only depends on `dp[j]` for indices `j` with smaller values. This ordering is what makes the DP well-defined and lets you memoize safely.

### Hint 3

The blocking rule ("`arr[i] > arr[k]` for **all** `k` between `i` and `j`") is the key simplification. When you scan outward from `i` (up to `d` steps left and `d` steps right), the **first** bar you meet that is `>= arr[i]` blocks everything beyond it in that direction — so you can simply **break out of the scan** the moment you hit such a bar. You never need to look further. This turns each index's transition into two short linear scans of length at most `d`.

## Approach

Define `dp[i]` = the maximum number of indices you can visit on a path that **starts** at index `i` (counting `i` itself). The answer is `max(dp[i])` over all `i`.

Transition: from `i` you may jump to `j = i ± x` for `1 <= x <= d`, but only if `arr[i] > arr[j]` and every bar strictly between `i` and `j` is also shorter than `arr[i]`. Because of the "all bars between" condition, if you scan outward from `i` and encounter any bar `>= arr[i]`, that bar (and everything past it in that direction) is unreachable — so you stop scanning that direction immediately.

So for each `i`:

```
dp[i] = 1 + max( dp[j] for every reachable neighbor j ),  or 1 if none.
```

We compute this with **memoized DFS**. Because every jump strictly decreases the bar height, recursion always descends to smaller values and terminates; each `dp[i]` is computed once and cached. (Equivalently, you could sort indices by height ascending and fill `dp` bottom-up — when you process `i`, all reachable `j` have smaller heights and are already done.)

Walking the scan from index `i` to the right:

```
for x = 1 .. d, j = i + x (in bounds):
    if arr[j] >= arr[i]: break        # blocked, stop scanning right
    dp[i] = max(dp[i], 1 + dfs(j))
```

and symmetrically to the left. Finally take the max over all `dp[i]`.

**Example:** `arr = [6,4,14,6,8,13,9,7,10,6,12], d = 2`. Starting at index 10 (value 12), you can reach index 8 (value 10, the bar at index 9 = 6 is shorter so it does not block), then from 8 reach index 6 (value 9), then from 6 reach index 7 (value 7). That path `10 -> 8 -> 6 -> 7` visits 4 indices, which is the maximum.

This is a genuinely Hard problem — the jump rules take care to read precisely — but once you see that "first taller-or-equal bar blocks the rest," the implementation is short and the DP is clean. Take it one direction at a time.

## Complexity Analysis

Time Complexity: O(n * d) — each of the `n` indices does two scans of length at most `d`, and memoization ensures each `dp[i]` is computed only once. (A bottom-up version adds an O(n log n) sort, which is dominated when `d` is large.)
Space Complexity: O(n) for the memo array plus up to O(n) recursion depth in the worst case (a fully descending staircase).

## Edge Cases

- **Single element** (`arr = [5]`): no jumps possible, the answer is `1`. The DP base case must count the starting index itself.
- **All equal values** (`arr = [3,3,3,3,3]`): no jump is ever legal because `arr[i] > arr[j]` never holds (it must be strictly greater). Every `dp[i] = 1`, so the answer is `1`. Watch for using `>=` instead of strict `>`.
- **Strictly decreasing within reach** (`arr = [7,6,5,4,3,2,1], d = 1`): you can chain through every index, so the answer equals the array length. Confirms the chaining/transition is correct.
- **A bar blocking a closer-but-shorter bar**: e.g. you could jump 2 steps to a short bar, but a tall bar sits in between — the tall bar blocks it. The `break`-on-`>=` logic handles this; a naive "check each `j` independently" without breaking would wrongly allow the jump.
- **`d` larger than the meaningful range**: scans are bounded by both `d` and the array edges, so large `d` is harmless — the `break` stops the scan early at the first taller bar anyway.
