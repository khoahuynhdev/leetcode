# 1658. Minimum Operations to Reduce X to Zero

[LeetCode Link](https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/)

Difficulty: Medium
Topics: Array, Hash Table, Binary Search, Sliding Window, Prefix Sum
Acceptance Rate: 41.7%

## Hints

### Hint 1

The operations only ever peel elements off the two ends, so whatever you remove is always a **prefix plus a suffix** of the original array — never anything from the middle. That means the order of your removals doesn't matter, only *how many* you take from the left and *how many* from the right. Try to restate the problem in terms of those two counts instead of a sequence of moves.

Also notice what's missing from the problem: nothing asks you to actually mutate the array. Simulating removals (or exploring both choices recursively) is the trap that makes this look harder than it is.

### Hint 2

Searching directly over "prefix of length `i` + suffix of length `j` with sum exactly `x`" is doable with prefix sums and a two-pointer scan, but there's a slicker reformulation. The elements you **keep** are exactly the ones you didn't remove — and since the removals are a prefix and a suffix, the kept elements form a **contiguous subarray** in the middle.

So: if `total` is the sum of the whole array, and the removed part sums to `x`, what must the middle subarray sum to? And if you want to *minimize* the number removed, what do you want to do to the middle?

### Hint 3

Minimize removals ⟺ **maximize the length of a contiguous subarray whose sum equals `total - x`**. Call that target `target = total - x`. If you find the longest such subarray with length `L`, the answer is `n - L`.

The last piece is why a sliding window is legal here: the constraint `1 <= nums[i]` means every element is strictly positive, so the running window sum is **monotonically increasing** as you extend the right end and **decreasing** as you advance the left end. That monotonicity is exactly what lets you move `left` forward without ever needing to back up — giving an O(n) scan instead of a hash-map-of-prefix-sums lookup. (The hash map version works too and is the approach you'd need if negatives were allowed.)

Watch the two degenerate targets: `target < 0` and `target == 0`.

## Approach

**Step 1 — Reframe.** Removing a prefix summing to `p` and a suffix summing to `s` with `p + s = x` leaves behind a contiguous middle subarray summing to `total - x`. Minimizing `(prefix length + suffix length)` is the same as maximizing the middle subarray's length. So the whole problem collapses to:

> Find the **longest** contiguous subarray of `nums` with sum exactly `target = total - x`.

If such a subarray of length `L` exists, the answer is `n - L`. If none exists, the answer is `-1`.

**Step 2 — Handle the degenerate targets.**

- `target < 0`: the entire array sums to less than `x`, so even removing everything can't reach `x`. Return `-1` immediately. (This also guards the sliding window, which assumes a non-negative target.)
- `target == 0`: you must remove the entire array. The "longest subarray summing to 0" is the empty one, length `0`, so the answer is `n`. Returning `n` up front keeps the main loop's `sum == target` check dealing only with non-empty windows.

**Step 3 — Slide the window.** Maintain `left`, `right`, and the running `sum` of `nums[left..right]`.

- Extend: add `nums[right]` to `sum`.
- Shrink: while `sum > target`, subtract `nums[left]` and advance `left`. Since every element is `>= 1`, each subtraction strictly decreases `sum`, so this terminates, and `sum` can never skip past `target` from below — we check it right after shrinking.
- Record: if `sum == target`, update the best length with `right - left + 1`.

Each index enters and leaves the window at most once, so the pointers advance a total of at most `2n` times.

**Worked example** — `nums = [3,2,20,1,1,3]`, `x = 10`:

`total = 30`, so `target = 30 - 10 = 20`.

| right | value | sum before shrink | shrink action | window | sum | best |
|---|---|---|---|---|---|---|
| 0 | 3 | 3 | — | `[3]` | 3 | — |
| 1 | 2 | 5 | — | `[3,2]` | 5 | — |
| 2 | 20 | 25 | drop 3, drop 2 | `[20]` | 20 | **1** |
| 3 | 1 | 21 | drop 20 | `[1]` | 1 | 1 |
| 4 | 1 | 2 | — | `[1,1]` | 2 | 1 |
| 5 | 3 | 5 | — | `[1,1,3]` | 5 | 1 |

Longest window is `[20]` with `L = 1`, so the answer is `6 - 1 = 5` — matching the expected output (remove `3,2` from the left and `1,1,3` from the right).

**Why not simulate?** Greedily taking the smaller end fails (Example 3 forces you to remove a `3` early even though `1`s sit on the right), and branching on both ends is `O(2^n)`. Memoizing on `(left, right)` is `O(n^2)` states — too slow for `n = 10^5`. The reframing is what buys the linear scan.

## Complexity Analysis

Time Complexity: O(n) — one pass to compute `total`, then one pass where `right` advances `n` times and `left` advances at most `n` times total.
Space Complexity: O(1) — only a handful of integer accumulators; the input is never copied or mutated.

## Edge Cases

- **`x > total` (`target < 0`)**: impossible — removing every element still leaves `x` positive. Must return `-1` *before* entering the window loop, otherwise `sum > target` would be true even for an empty window and `left` would run past `right` forever chasing an unreachable target.
- **`x == total` (`target == 0`)**: the answer is `n`, not `-1`. Easy to lose if you only look for non-empty windows, since the matching "middle" subarray here is the empty one.
- **No subarray sums to `target` even though `target > 0`**: e.g. `nums = [2,2,2], x = 3` → `target = 3`, but every subarray sum is even. The window never hits equality, so `best` stays unset and the answer is `-1`. Don't confuse "no window found" with "window of length 0".
- **Single-element array**: `[5], x = 5` → `1`; `[5], x = 3` → `-1`. Confirms the loop handles `n = 1` without special casing.
- **Answer uses only one end**: `nums = [1,1,4,2,3], x = 5` takes purely a suffix. The reformulation covers this automatically because the middle subarray is simply flush against one edge — no separate "prefix-only / suffix-only" branch needed.
- **Integer sizes**: `n <= 10^5` and `nums[i] <= 10^4` put `total` at up to `10^9`, and `x <= 10^9`, so `total - x` fits comfortably in Go's `int` (64-bit on all mainstream platforms). No overflow concern here, but it's worth a moment's thought since the sums are near the 32-bit boundary.
- **All elements positive is load-bearing**: the sliding window is only correct because of it. If negatives were allowed you'd need the prefix-sum + hash-map variant (store the *earliest* index for each prefix sum, then maximize `i - firstIndex[prefix-target]`), which is the Hash Table topic tag hinting at the general version.
