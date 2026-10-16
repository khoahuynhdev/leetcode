# 2916. Subarrays Distinct Element Sum of Squares II

[LeetCode Link](https://leetcode.com/problems/subarrays-distinct-element-sum-of-squares-ii/)

Difficulty: Hard
Topics: Array, Hash Table, Binary Indexed Tree, Segment Tree
Acceptance Rate: 38.9%

## Hints

### Hint 1

With n up to 10^5 there are about 5·10^9 subarrays, so you can't look at each one. Try fixing the **right endpoint** `r` and thinking about the distinct counts of all subarrays that end at `r` together, as one array indexed by the left endpoint `l`.

### Hint 2

Let `f(l)` be the distinct count of `nums[l..r]`. When you move from `r-1` to `r`, which `f(l)` values change? Only the subarrays that didn't already contain `nums[r]` get a new distinct value. Those are exactly the left endpoints after the previous occurrence of `nums[r]`. So each step is a **range +1 update** on a contiguous block of `f`.

### Hint 3

You need the sum of `f(l)^2`, but adding 1 to `x` changes `x^2` by `2x + 1`. So adding 1 to every `f(l)` in a range `[a, b]` raises the sum of squares by `2 · sum(f[a..b]) + (b - a + 1)`. That means you only need **range add + range sum**, which a lazy segment tree handles in O(log n).

## Approach

1. Sweep `r` from `0` to `n-1`. Keep a map `last` from value to the index where it last appeared.
2. For the current `r`, let `p = last[nums[r]]` (or `-1` if it hasn't appeared yet). Every left endpoint `l` in `[p+1, r]` gains a new distinct element, so `f(l)` goes up by 1. Left endpoints `l <= p` already contain `nums[r]` and stay the same. (`f(r)` starts at 0 and becomes 1, which this same update covers.)
3. Keep a running value `squares = Σ f(l)^2` over all `l <= r`. Before the update, query `s = sum(f[p+1..r])` from the segment tree. Then:
   - `squares += 2·s + (r - p)`
   - apply range `+1` on `[p+1, r]` in the tree.
4. `squares` is now the sum of squared distinct counts of all subarrays ending at `r`. Add it to the answer (mod 10^9+7).

**Example** `nums = [1,2,1]`:

| r | range updated | f after update | squares | answer |
|---|---------------|----------------|---------|--------|
| 0 | [0,0]         | [1]            | 1       | 1      |
| 1 | [0,1]         | [2,1]          | 5       | 6      |
| 2 | [1,2] (prev 1 at 0) | [2,2,1]  | 9       | 15     |

**Why it works:** the segment tree always holds the exact values of `f`, and the identity `(x+1)^2 = x^2 + 2x + 1` lets us update the sum of squares from just the range sum, without storing the squares themselves. Raw sums of `f` are at most n² ≈ 10^10, so they fit in `int64` without taking the modulus. Only `squares` and the answer need it.

## Complexity Analysis

Time Complexity: O(n log n): one range query and one range update per index.
Space Complexity: O(n) for the segment tree and the last-occurrence map.

## Edge Cases

- **Single element**: one subarray with distinct count 1, so the answer is 1. Check that the tree works when n = 1.
- **All identical values**: every update covers only `[r, r]`, so each `f(l)` stays at 1. This tests that the previous-occurrence boundary is handled correctly.
- **All distinct values**: every update covers the whole prefix `[0, r]`, so this gives the largest totals. The answer must be reduced mod 10^9+7, and intermediate products must not overflow.
- **Value reappearing after a gap** (e.g. `[1,2,1]`): only left endpoints after the previous occurrence should be incremented. An off-by-one here (`p` vs `p+1`) changes the result.
