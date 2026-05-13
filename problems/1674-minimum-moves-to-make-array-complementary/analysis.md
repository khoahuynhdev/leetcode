# 1674. Minimum Moves to Make Array Complementary

[LeetCode Link](https://leetcode.com/problems/minimum-moves-to-make-array-complementary/)

Difficulty: Medium
Topics: Array, Hash Table, Prefix Sum
Acceptance Rate: 51.7%

## Hints

### Hint 1

The array is made of `n/2` independent pairs `(nums[i], nums[n-1-i])`. Every pair must end up summing to the *same* target `T`. Since `1 <= nums[i] <= limit`, the only possible targets are `2 .. 2*limit`. Think about how much each pair would "cost" for a fixed `T`, and then about how to evaluate all targets efficiently instead of one at a time.

### Hint 2

For a single pair with values `a` and `b` (let `lo = min(a,b)`, `hi = max(a,b)`), and a chosen target `T`:
- `0` moves if `a + b == T`.
- `1` move if `T` is reachable by changing exactly one element, i.e. `lo + 1 <= T <= hi + limit`.
- `2` moves otherwise.

So as `T` sweeps from `2` to `2*limit`, each pair's cost is a step function: it starts at `2`, drops to `1` on the interval `[lo+1, hi+limit]`, and dips to `0` only at the single point `lo+hi`. You need the sum of all these step functions, then the minimum over `T`.

### Hint 3

Range-add + point-update over a 1D axis screams **difference array** (a prefix-sum trick). Start every target with a baseline cost of `2 * (n/2) = n`. For each pair, apply `-1` over the range `[lo+1, hi+limit]` (the "one move suffices" region) and an additional `-1` at the single point `lo+hi` (the "zero moves" region, which is nested inside the one-move region). After processing all pairs, take a running prefix sum across `T = 2 .. 2*limit` and track the minimum. That's `O(n + limit)`.

## Approach

Pair up the array: pair `k` is `(nums[k], nums[n-1-k])` for `k = 0 .. n/2 - 1`. These pairs are completely independent — choosing the common target `T` is the only shared decision.

For a fixed target `T` and a pair with sorted values `lo <= hi`:

1. If `lo + hi == T`, no change is needed → cost `0`.
2. Otherwise, can one change fix it? We need the other (unchanged) element `x` to satisfy `T - x` being a legal value in `[1, limit]`, i.e. `1 <= T - x <= limit`. Taking the best of the two elements, this is possible exactly when `lo + 1 <= T <= hi + limit` → cost `1`.
3. Otherwise both elements must change → cost `2` (always achievable since `T` itself lies in `[2, 2*limit]`).

Now we want, for each `T` in `[2, 2*limit]`, the total cost summed over all pairs, and finally the minimum.

Computing this naively is `O(limit * n)`. Instead, use a difference array `delta` of length `2*limit + 2`:

- Conceptually every pair contributes `2` for every target, so the baseline total is `n` (that's `2 * n/2`).
- For each pair: it actually costs `1` (not `2`) on `[lo+1, hi+limit]`, so add `-1` to that range: `delta[lo+1] -= 1; delta[hi+limit+1] += 1`.
- And it costs `0` (one *further* less than `1`) exactly at `T = lo+hi`, so add another `-1` there: `delta[lo+hi] -= 1; delta[lo+hi+1] += 1`.

Then walk `T` from `2` to `2*limit`, maintaining `cur = n + (running sum of delta up to T)`, and keep the minimum `cur`. That minimum is the answer.

### Worked example

`nums = [1,2,4,3]`, `limit = 4`. Pairs: `(1,3)` → `lo=1, hi=3`; `(2,4)` → `lo=2, hi=4`. Baseline total = `4`.

- Pair `(1,3)`: one-move range `[2, 7]` → `delta[2]-=1, delta[8]+=1`; zero-move point `4` → `delta[4]-=1, delta[5]+=1`.
- Pair `(2,4)`: one-move range `[3, 8]` → `delta[3]-=1, delta[9]+=1`; zero-move point `6` → `delta[6]-=1, delta[7]+=1`.

Sweeping: `T=2 → 3`, `T=3 → 2`, `T=4 → 1`, `T=5 → 2`, `T=6 → 1`, `T=7 → 2`, `T=8 → 3`. Minimum is `1`. ✔

This is a neat problem once the step-function picture clicks, but getting the `1`-move bound `[lo+1, hi+limit]` exactly right (and remembering the `0`-move point is *nested* inside it, so it needs an extra `-1`) is the part most people stumble on. Take your time with that interval.

## Complexity Analysis

Time Complexity: O(n + limit) — one pass over the `n/2` pairs to fill the difference array, then one pass over the `~2*limit` possible targets.
Space Complexity: O(limit) — the difference array has size `2*limit + 2`.

## Edge Cases

- **`n == 2` (single pair):** Any single pair is trivially complementary on its own (there's only one sum to match), so the answer is always `0`. The algorithm handles this naturally — the zero-move point of that pair always yields a total of `0`.
- **Already complementary array:** Every pair shares the same sum `S`, so at `T = S` every pair costs `0` and the minimum is `0`.
- **`limit` very small (e.g. `limit == 1`):** The one-move range `[lo+1, hi+limit]` can collapse to a single point or even be empty-feeling; indexing must stay within the `2*limit + 2` array. Check that `hi + limit + 1 <= 2*limit + 1`, which holds since `hi <= limit`.
- **All elements equal but far from a feasible shared sum given a tiny `limit`:** e.g. `nums = [1,2,2,1], limit = 2` forces `2` moves on some pair — the answer is `2`, not `1`. Verifies the "two moves" fallback.
- **Off-by-one on array bounds:** `delta` must have length `2*limit + 2` so that `delta[2*limit + 1]` (from `hi + limit + 1` when `hi == limit`, and from `lo + hi + 1` when `lo + hi == 2*limit`) is valid. Loop targets only over `2 .. 2*limit` inclusive.
