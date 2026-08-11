# 2996. Smallest Missing Integer Greater Than Sequential Prefix Sum

[LeetCode Link](https://leetcode.com/problems/smallest-missing-integer-greater-than-sequential-prefix-sum/)

Difficulty: Easy
Topics: Array, Hash Table, Sorting
Acceptance Rate: 37.7%

## Hints

### Hint 1

The problem is really two small, independent tasks glued together: first compute a number, then search for a number. Solve them one at a time instead of trying to do both in a single clever pass. For the second half, ask yourself what data structure answers "is this value in the array?" in constant time.

### Hint 2

The word *prefix* is doing a lot of work here. You only ever look at runs that start at index `0` — a longer consecutive run sitting in the middle of the array is irrelevant. So walk forward from index `0`, accumulating the sum, and stop the instant `nums[j] != nums[j-1]+1`. There is no scanning to restart, no maximum to track across multiple candidates.

### Hint 3

Once you have the sum `s`, the answer is just the first integer `x >= s` that isn't in `nums`. Put every element in a hash set and increment `x` from `s` until it misses. This terminates fast: values are capped at 50, so at most 51 distinct values can block you — after at most 51 steps you must land on a gap. No sorting or binary search is needed.

## Approach

**Step 1 — find the longest sequential prefix sum.**

By definition, the prefix `nums[0..0]` is always sequential, so start with `sum = nums[0]` and `i = 1`. Advance while `nums[i] == nums[i-1] + 1`, adding each element to `sum`. The first violation ends the prefix permanently — because a prefix must begin at index `0`, a break at index `i` means no prefix of length `> i` can be sequential. This is why a single forward scan with an early exit is exact, not a heuristic.

For `nums = [3,4,5,1,12,14,13]`: `sum = 3`, then `4 == 3+1` so `sum = 7`, then `5 == 4+1` so `sum = 12`, then `1 != 5+1` so we stop. The longest sequential prefix is `[3,4,5]` with sum `12`. Note that `[12,13,14]` also forms a consecutive run, but it isn't a prefix, so it doesn't count.

**Step 2 — find the smallest missing integer at least `sum`.**

Build a `map[int]struct{}` (or a fixed-size boolean array, since values are bounded by 50) holding every element of `nums`. Then set `x = sum` and increment while `x` is present in the set. The first `x` not in the set is the answer.

Continuing the example: `x = 12` is present, `13` is present, `14` is present, `15` is not — return `15`.

Two details worth internalizing:

- The comparison is `>=`, not `>`. If `sum` itself is missing from `nums`, `sum` is the answer immediately. In example 1, `sum = 6` and `6` is absent, so the loop never runs.
- The scan is guaranteed to terminate quickly. `nums` holds at most 50 elements with values in `[1, 50]`, so at most 50 consecutive integers can be blocked. Once `x` exceeds 50 nothing can be in the set, so the loop runs at most 51 iterations regardless of how large `sum` is (the max possible sum is `1 + 2 + ... + 50 = 1275`).

The "Sorting" tag on this problem hints at an alternative — sort the array and walk it — but the hash-set approach is simpler and strictly faster here.

## Complexity Analysis

Time Complexity: O(n) — one pass to compute the prefix sum, one pass to build the set, and at most 51 (i.e. O(1)) probes for the missing value. With `n <= 50` this is trivially fast.

Space Complexity: O(n) for the hash set. Using a fixed `[51]bool` lookup table instead makes this O(1) auxiliary space, since values are bounded by the constraints.

## Edge Cases

- **Single element, e.g. `[1]`.** The prefix consisting of just `nums[0]` is sequential by definition, so `sum = 1`. Since `1` is in the array, the answer is `2`. A loop written as `for i := 1; ...` handles this naturally, but an implementation that requires at least two elements to seed the sum will break here.
- **The very first pair breaks the run, e.g. `[5,1,2,3]`.** The longest sequential prefix is just `[5]`, sum `5`. It's tempting to grab the longer run `[1,2,3]` later in the array — but that isn't a prefix. Answer: `5` is present, so `6`.
- **`sum` itself is missing, e.g. `[1,2,3,2,5]`.** Sum is `6`, which isn't in `nums`, so the answer is `6` with zero probing. This is the `>=` vs `>` trap.
- **The answer exceeds every array value, e.g. `[50]`.** Sum is `50`, which is present, so the answer is `51` — larger than any element. Don't cap the search at `max(nums)` or at the array length.
- **Entire array is sequential, e.g. `[1,2,3,4,5]`.** The scan must run to the end without indexing out of bounds. Sum is `15`, which is absent, so the answer is `15`.
- **Duplicate values after the prefix, e.g. `[1,2,1,2,3,4]`.** The prefix is `[1,2]` with sum `3`. Duplicates are harmless for a set-based membership check: `3` and `4` are present, so the answer is `5`.
