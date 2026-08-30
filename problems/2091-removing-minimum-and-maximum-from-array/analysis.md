# 2091. Removing Minimum and Maximum From Array

[LeetCode Link](https://leetcode.com/problems/removing-minimum-and-maximum-from-array/)

Difficulty: Medium
Topics: Array, Greedy
Acceptance Rate: 61.1%

## Hints

### Hint 1

The values in the array are mostly a distraction. Only two of them ever matter: the minimum and the maximum. Once you know *where* those two live, the rest of the array is just padding. Try restating the problem purely in terms of two indices in a length-`n` array.

### Hint 2

Deletions only happen at the two ends, so any sequence of deletions leaves behind a contiguous middle segment. That means the whole strategy is described by "how many did I take from the front" and "how many did I take from the back" — the *order* of the deletions never matters. Both target elements must end up outside the surviving middle segment.

### Hint 3

With two indices `lo <= hi`, there are exactly three ways to get both of them out, and you just take the cheapest:

1. Sweep from the front past both: `hi + 1` deletions.
2. Sweep from the back past both: `n - lo` deletions.
3. Meet in the middle — take `lo` from the front and `hi` from the back: `(lo + 1) + (n - hi)` deletions.

There is no fourth option, because a strategy that reaches `lo` from the back has already passed `hi`, and vice versa.

## Approach

The key reframing is that this is not really an array problem, it's a two-index geometry problem.

**Step 1 — locate the two targets.** Scan the array once, tracking the index of the smallest value and the index of the largest value. The problem guarantees distinct integers, so both are unambiguous. Let `i` be the min's index and `j` the max's index.

**Step 2 — normalize.** We don't care which one is the min and which is the max, only their positions. Set `lo = min(i, j)` and `hi = max(i, j)`. Let `n = len(nums)`.

**Step 3 — enumerate the three strategies.** Every deletion removes an element from the front or the back, so after any number of deletions the array that remains is a contiguous slice `nums[a : n-b]` where `a` is the front count and `b` is the back count, costing `a + b`. We need both `lo` and `hi` excluded from that slice. Since `lo <= hi`, exactly three minimal configurations exist:

- **Front only:** delete through index `hi`, which necessarily also removes `lo`. Cost `hi + 1`.
- **Back only:** delete back through index `lo`, which necessarily also removes `hi`. Cost `n - lo`.
- **Split:** remove `lo` from the front (cost `lo + 1`) and `hi` from the back (cost `n - hi`). Total `(lo + 1) + (n - hi)`.

Any other combination is dominated by one of these — e.g. deleting extra elements past `hi` from the front only adds cost.

**Step 4 —** return the minimum of the three.

**Worked example.** `nums = [2, 10, 7, 5, 4, 1, 8, 6]`, so `n = 8`. The minimum `1` sits at index 5, the maximum `10` at index 1, giving `lo = 1`, `hi = 5`.

- Front only: `5 + 1 = 6`
- Back only: `8 - 1 = 7`
- Split: `(1 + 1) + (8 - 5) = 2 + 3 = 5`

The answer is `5`, matching the expected output: two from the front (`2, 10`) and three from the back (`8, 6` plus `1`).

One nice property of the split formula is that it never *undercounts* when `lo == hi` (a single-element array). There it evaluates to `n + 1`, which loses to the other two options, so no special-casing is needed.

This is a Medium mostly because of the "three cases, and only three" reasoning — the code itself is a single pass and a couple of `min` calls. If you found yourself reaching for two pointers or a sliding window, that's a very common first instinct here; the trap is treating deletion as a process rather than as a choice of a surviving window.

## Complexity Analysis

Time Complexity: O(n) — one pass to find the two indices, then constant work.
Space Complexity: O(1) — only a handful of index variables.

## Edge Cases

- **Single element (`n == 1`):** the lone element is both the minimum and the maximum, so `lo == hi == 0` and the answer is `1`. The front-only and back-only formulas both yield `1`; the split formula yields `2` and is correctly discarded by the `min`.
- **Two elements:** whichever order they appear in, `lo = 0` and `hi = 1`, and all three formulas give `2`. You must delete both, so `2` is right.
- **Min and max at opposite ends:** e.g. `[1, 3, 4, 2, 5]`. The split strategy wins with `2`, while the sweep strategies would cost `n`. This is the case that makes the third formula necessary.
- **Min and max adjacent near one end:** e.g. `[0, -4, 19, 1, 8, -2, -3, 5]`. A single-sided sweep beats the split, which is why you can't just always meet in the middle.
- **Max appears before min:** the `lo`/`hi` normalization handles this; forgetting it and hardcoding `minIdx <= maxIdx` produces negative or inflated counts.
- **All negative values:** initializing the running min/max to `0` instead of `nums[0]` breaks on inputs like `[-5, -1, -3]`. Seed from the first element.
