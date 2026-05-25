# 1871. Jump Game VII

[LeetCode Link](https://leetcode.com/problems/jump-game-vii/)

Difficulty: Medium
Topics: String, Dynamic Programming, Sliding Window, Prefix Sum
Acceptance Rate: 28.8%

## Hints

### Hint 1

This is a reachability question: "can I get from index `0` to index `n-1`?" That instantly suggests dynamic programming where `dp[i]` means "index `i` is reachable." The interesting twist is that you don't jump a fixed distance — from `i` you may land anywhere in the window `[i+minJump, i+maxJump]`. Think about how you can reverse that relationship: which earlier indices could have jumped *to* `i`?

### Hint 2

Flip the direction of the jump. Index `i` is reachable if `s[i] == '0'` **and** at least one reachable index exists in the window `[i-maxJump, i-minJump]`. A naive DP would rescan that whole window for every `i`, which is `O(n * (maxJump-minJump))` and can degrade to `O(n^2)`. The window slides forward by exactly one as `i` increases — that's the classic cue for a sliding window or a prefix sum.

### Hint 3

The key insight: you don't need *which* earlier index is reachable, only *whether one exists* in the window. So maintain a prefix-sum (a running count) of reachable indices. Then "is there a reachable index in `[lo, hi]`?" becomes the `O(1)` query `pre[hi] - pre[lo-1] > 0`. Computing every `dp[i]` this way turns the whole algorithm into a single `O(n)` pass.

## Approach

This problem looks intimidating because of the variable jump range, but reversing the jump makes it clean.

Define `dp[i] = true` when index `i` is reachable from index `0`. Index `0` is reachable by definition (and the constraints guarantee `s[0] == '0'`), so `dp[0] = true`.

For any later index `i`, you could only have arrived from some index `j` that jumped forward into `i`. The jump rule `j + minJump <= i <= j + maxJump` rearranges to:

```
i - maxJump <= j <= i - minJump
```

So `dp[i]` is `true` exactly when both:

1. `s[i] == '0'` (you are allowed to land on `i`), and
2. some index `j` in the window `[i-maxJump, i-minJump]` has `dp[j] == true`.

Checking condition 2 by scanning the window each time is too slow in the worst case. Instead, keep a prefix-sum array `pre`, where `pre[i]` counts how many of `dp[0..i]` are `true`. The number of reachable indices in `[lo, hi]` is `pre[hi] - pre[lo-1]` (treating `pre[-1]` as `0`). If that count is positive, at least one reachable predecessor exists, so `dp[i] = true`.

Walk `i` from `1` to `n-1`:

- Clamp `lo = max(0, i-maxJump)` and `hi = i-minJump`.
- If `hi >= 0` and `s[i] == '0'`, set `dp[i]` based on the prefix-sum query over `[lo, hi]`.
- Update `pre[i] = pre[i-1] + (dp[i] ? 1 : 0)`.

The answer is `dp[n-1]`.

**Worked example:** `s = "011010"`, `minJump = 2`, `maxJump = 3`.

- `dp[0] = true`.
- `i = 3` (`s[3] = '0'`): window `[0, 1]`, and `dp[0]` is reachable → `dp[3] = true`.
- `i = 5` (`s[5] = '0'`): window `[2, 3]`, and `dp[3]` is reachable → `dp[5] = true`.
- `dp[5]` is the last index → answer `true`.

Note the same prefix-sum idea can be implemented as a sliding count that adds `dp[i-minJump]` and removes `dp[i-maxJump-1]` as the window advances; the prefix-sum version is used here because it reads more directly.

## Complexity Analysis

Time Complexity: O(n) — a single pass over the string, with `O(1)` work per index thanks to the prefix-sum query.
Space Complexity: O(n) — the `dp` and prefix-sum arrays. (This can be reduced to O(n) for `dp` alone or O(1) extra with a rolling window count, but O(n) is the clear formulation.)

## Edge Cases

- **Last character is `'1'`:** index `n-1` can never be landed on, so the answer must be `false`. The `s[i] == '0'` check handles this automatically.
- **Minimal length `n == 2`:** with `minJump == maxJump == 1`, a single jump from `0` to `1` decides the answer; make sure the window math (`hi = i - minJump = 0`) still queries index `0`.
- **`hi < 0` (i.e. `i < minJump`):** these early indices are simply unreachable because no valid predecessor window exists yet — skip them without touching the prefix sum incorrectly.
- **Window clamped at the left (`i - maxJump < 0`):** clamp `lo` to `0` so the prefix-sum query stays in bounds; reachable indices are never negative anyway.
- **`pre[lo-1]` when `lo == 0`:** treat the missing prefix as `0` to avoid an out-of-bounds read.
- **Blocked corridor:** a run of `'1'`s longer than `maxJump - minJump` can sever the path even when the endpoints are `'0'` (this is exactly why Example 2 returns `false`).
