# 3471. Find the Largest Almost Missing Integer

[LeetCode Link](https://leetcode.com/problems/find-the-largest-almost-missing-integer/)

Difficulty: Easy
Topics: Array, Hash Table
Acceptance Rate: 38.9%

## Hints

### Hint 1

The constraints are tiny (`n <= 50`), so a brute-force sweep over every window is a legitimate accepted answer. But before you write it, notice that the question isn't really "how many times does `x` occur?" — it's "how many *windows* does `x` land in?" Those are different numbers, and separating them is the whole problem. A frequency map plus a little thinking about window geometry will get you to something much faster than brute force.

### Hint 2

Forget values for a moment and think about a single **index** `i`. Exactly which windows of size `k` cover it? The windows are indexed by their start `s`, and `i` is covered when `s <= i <= s + k - 1`, i.e. `s` ranges over `[max(0, i-k+1), min(i, n-k)]`. Count how many starts that is. For most indices the count is comfortably larger than 1 — try it on a concrete array like `[3,9,2,1,7]` with `k = 3` and see which indices are the exceptions.

### Hint 3

Work the count out: index `i` is covered by `min(i, n-k) - max(0, i-k+1) + 1` windows. When `1 < k < n`, that expression equals 1 **only** at `i = 0` and `i = n-1`; every interior index is covered at least twice. So only the first and last elements can possibly be almost missing — and each qualifies only if it appears *nowhere else* in the array (a second occurrence would drag in a second window, and if the value sits at both ends those are two distinct windows). That leaves just two special cases to reason about separately: `k == n` (a single window, so every value in it is almost missing) and `k == 1` (each window is one element, so a value is almost missing exactly when it occurs once).

## Approach

The key reframe is that "appears in exactly one subarray of size `k`" is a statement about **window coverage**, not about occurrence count. A value `x` is almost missing when the *union* of the window-sets of all its occurrences has size exactly 1.

**Step 1 — count coverage per index.** Windows of size `k` are identified by their start `s ∈ [0, n-k]`. A window covers index `i` iff `s <= i <= s+k-1`, so the valid starts are `s ∈ [max(0, i-k+1), min(i, n-k)]`, giving

```
cover(i) = min(i, n-k) - max(0, i-k+1) + 1
```

**Step 2 — find where `cover(i) == 1`.** Assume `1 < k < n` and split on the two `min`/`max` branches:

| branch | `cover(i)` | when is it 1? |
| --- | --- | --- |
| `i <= k-1`, `i <= n-k` | `i + 1` | `i == 0` |
| `i <= k-1`, `i > n-k` | `n - k + 1` | never (`k < n`) |
| `i > k-1`, `i <= n-k` | `k` | never (`k > 1`) |
| `i > k-1`, `i > n-k` | `n - i` | `i == n-1` |

So only the **endpoints** `0` and `n-1` sit in a single window.

**Step 3 — lift from indices to values.** An endpoint's value qualifies only if it has no other occurrence. If `x` also appears at some interior index `j`, then `j` is covered by ≥ 2 windows, so `x`'s union is ≥ 2. If `x` sits at *both* ends, its union is `{window 0, window n-k}` — two distinct windows since `k < n`. Either way, the endpoint value must have frequency exactly 1. A hash map of value → count answers that in O(1) per lookup.

**Step 4 — handle the two degenerate cases.**

- `k == n`: there is exactly one window and it covers everything, so every distinct value is almost missing. Return `max(nums)`.
- `k == 1`: every window is a single element, so `x`'s window count equals its occurrence count. Return the largest value with frequency 1, or `-1`.

(When `n == 1` both cases apply and both return `nums[0]`, so the ordering of the checks doesn't matter.)

**Walking Example 1** — `nums = [3,9,2,1,7], k = 3`, so `n = 5` and `1 < k < n`. Frequencies are all 1. Candidate `nums[0] = 3` has frequency 1 → valid. Candidate `nums[4] = 7` has frequency 1 → valid. The answer is `max(3, 7) = 7`, matching the exhaustive count in the statement.

**Walking Example 2** — `nums = [3,9,7,2,1,7], k = 4`, `n = 6`. Candidate `nums[0] = 3` has frequency 1 → valid. Candidate `nums[5] = 7` has frequency 2 (it also sits at index 2) → rejected, and indeed index 2 is covered by all three windows. Answer: `3`.

Since values are guaranteed non-negative, `-1` is a safe sentinel for "nothing found" and needs no extra `found` flag.

## Complexity Analysis

Time Complexity: O(n) — one pass to build the frequency map, then a constant number of lookups (or one more pass over the map in the `k == 1` / `k == n` branches).
Space Complexity: O(n) for the frequency map. This could be reduced to O(1) extra space, since `0 <= nums[i] <= 50` allows a fixed 51-slot array instead of a map.

The straightforward brute force — enumerate all `n-k+1` windows, collect the distinct values in each, and tally — is O(n·k) time and passes easily at `n <= 50`. It's a perfectly good first submission; the O(n) version is the version worth understanding.

## Edge Cases

- **`k == n`** — a single window covering the entire array. Every distinct value appears in exactly one window, so the answer is `max(nums)`, never `-1`. Easy to get wrong if you jump straight to the "check both endpoints" rule, which would ignore the interior.
- **`k == 1`** — the endpoint rule collapses: every index is its own window, so the criterion becomes plain frequency-1, and interior elements are valid candidates. This is Example 3 (`[0,0], k=1` → `-1`).
- **`n == 1`** — forced `k == 1 == n`; both special cases fire and both correctly return `nums[0]`.
- **A repeated endpoint value** — e.g. `[5,3,5], k=2`: `nums[0]` and `nums[n-1]` are both `5`, but `5`'s two occurrences pull in two different windows, so the answer is `-1`. Checking frequency (not just "is it an endpoint") is what catches this.
- **Both endpoints valid, first one larger** — e.g. `[9,2,3,1], k=2` → `9`. Don't stop at the first qualifying candidate; take the max of the two.
- **No almost-missing integer at all** — e.g. `[1,2,3,1], k=2`. Return `-1`, and make sure the sentinel can't be confused with a real answer (safe here because `nums[i] >= 0`).
- **Zeros in the array** — `0` is a legal value, so a solution that uses `0` as "unset" or relies on truthiness will misreport. This is exactly why the sentinel is `-1`.
